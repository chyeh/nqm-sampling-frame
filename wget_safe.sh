#!/bin/bash

TIMEOUT=300
RETRYTIMES=3
DEVNULL=/dev/null

while getopts "H?q?s:t:T:R:p:O:C:P?" Option
do
	case $Option in
	H)
		[ -z "$HELPPARAM" ] && HELPPARAM=TRUE
	;;
	q)
		HIND_Q=-q
	;;
	s)
		SOURCE=$OPTARG
	;;
	t)
		TARGET=$OPTARG
	;;
	T)
		TIMEOUT=$OPTARG
	;;
	R)
		RETRYTIMES=$OPTARG
	;;
	p)
		PRIVILEGE=$OPTARG
	;;
	O)
		[ ! -z "$RUNMODE" ] && HELPPARAM=FALSE
		OUTPUT=$OPTARG
		RUNMODE=O
	;;
	C)
		[ ! -z "$RUNMODE" ] && HELPPARAM=FALSE
		CHDIR=$OPTARG
		RUNMODE=C
	;;
	P)
		[ ! -z "$RUNMODE" ] && HELPPARAM=FALSE
		PHYDIR=TRUE
		RUNMODE=P
	;;
	esac
done

help()
{
	echo "usage: wget_safe.sh OPTIONS"
	echo "OPTIONS:"
	echo "	-H              :	help"
	echo "	-q              :	quite mode"
	echo "	-s <url>        :	SOURCE File, full remote url"
	echo "	-t <localfile>  :	TARGET File, full local path"
	echo "	-T <secs>       :	TIMEOUT"
	echo "	-R <times>      :	RETRYTIMES"
	echo "	-p <owner>      :	PRIVILEGE"
	echo "	-O <outputfile> :	Output as file and check updated by diff"
	echo "	-C <outputdir>  :	Output directly"
	echo "	-P              :	Output directly to the physical path"
	echo "RETURN:"
	echo "	0 : normal, file updated"
	echo "	1 : normal, file noupdated"
	echo "	2 : exception, others"
	echo "	3 : exception, wgetError"
	echo "	4 : exception, tarError"
}

[ "$OPTIND" -le "$#" ] && help && exit 1

get_md5val()
{
	MD5FILE_=${1:-$DEVNULL}
	awk '{for(i=1; i<=NF;i++) if($i~"^[0-9a-z]+$" && length($i)==32) {print $i; break;}}' $MD5FILE_
}

wget_safe()
{
	if	wget $HIND_Q -T $TIMEOUT -Y off --tries=$RETRYTIMES $SOURCE.md5 -O $TARGET.md5.tmp$$
	then
		##if [ `cat $TARGET.md5.tmp$$ 2>$DEVNULL | wc -c` -eq 0 ]  ##���ص��յ�md5
		if [ ! -s $TARGET.md5.tmp$$ ]  ##���ص��յ�md5
		then
			rm -f $TARGET.md5.tmp$$ 1>$DEVNULL 2>&1
			return 3	#	fail on wget
		else
		{
			rm -f $TARGET.md5.last 1>$DEVNULL 2>&1
			mv -f $TARGET.md5 $TARGET.md5.last 1>$DEVNULL 2>&1
			mv -f $TARGET.md5.tmp$$ $TARGET.md5 1>$DEVNULL 2>&1

			if	[ ! -r $TARGET ] || [ "`get_md5val $TARGET.md5 2>$DEVNULL`" != "`get_md5val $TARGET.md5.last 2>$DEVNULL`" ]
			then
				if	wget $HIND_Q -T $TIMEOUT -Y off --tries=$RETRYTIMES $SOURCE -O $TARGET &&
					[ "`md5sum -b $TARGET | awk '{print $1}' 2>$DEVNULL`" == "`get_md5val $TARGET.md5 2>$DEVNULL`" ]
				then
                                       ## touch ../var/log/change.log
				       ## echo "$TARGET" >> ../var/log/change.log
					return 0	#	normal, updated
				else
					return 3	#	fail on wget
				fi
			else
				return 1		#	normal, noupdated
			fi
		}
		fi

	else
		rm -f $TARGET.md5.tmp$$ 1>$DEVNULL 2>&1
		return 3	#	fail on wget
	fi
}


wget_safe_onefile()
{
	wget_safe $SOURCE $TARGET
	RESULT=$?

	####echo $RESULT  $OUTPUT
	if	[ -f $TARGET ] && ([ "$RESULT" == "0" ] || [ ! -r $OUTPUT ])
	then
		if	tar jxfO $TARGET >$OUTPUT.new 2>$DEVNULL
		then

			if	[ ! -r $OUTPUT ] || ! diff -q $OUTPUT.new $OUTPUT 1>$DEVNULL 2>&1
			then
				cp -f $OUTPUT.new $OUTPUT
				[ ! -z $PRIVILEGE ] && chown -fRh $PRIVILEGE $TARGET $TARGET.md5 $OUTPUT 1>$DEVNULL 2>&1
				return 0
			else
				[ ! -z $PRIVILEGE ] && chown -fRh $PRIVILEGE $TARGET $TARGET.md5 $OUTPUT 1>$DEVNULL 2>&1
				return 1	#	normal, noupdated
			fi
		else
			[ ! -z $PRIVILEGE ] && chown -fRh $PRIVILEGE $TARGET $TARGET.md5 $OUTPUT 1>$DEVNULL 2>&1
			return 4	#	fail on tar
		fi
	else
		[ ! -z $PRIVILEGE ] && chown -fRh $PRIVILEGE $TARGET $TARGET.md5 $OUTPUT 1>$DEVNULL 2>&1
		return $RESULT
	fi
}

wget_safe_allfile()
{
	wget_safe $SOURCE $TARGET
	RESULT=$?

	if	[ "$RESULT" == "0" ]
	then
		[ ! -z $PRIVILEGE ] && chown -fRh $PRIVILEGE $TARGET $TARGET.md5 1>$DEVNULL 2>&1
		if	tar jxfC $TARGET $CHDIR
		then
			[ ! -z $PRIVILEGE ] && chown -fRh $PRIVILEGE $CHDIR 1>$DEVNULL 2>&1
			return 0
		else
			return 4
		fi
	else
		return $RESULT
	fi
}

wget_safe_allfile_phypath()
{
	wget_safe $SOURCE $TARGET
	RESULT=$?

	if	[ "$RESULT" == "0" ]
	then
		[ ! -z $PRIVILEGE ] && chown -fRh $PRIVILEGE $TARGET $TARGET.md5 1>$DEVNULL 2>&1
		if	tar jxfP $TARGET
		then
			[ ! -z $PRIVILEGE ] && tar jtf $TARGET |
			while read INNERFILENAME
			do
				chown -fRh $PRIVILEGE $INNERFILENAME 1>$DEVNULL 2>&1
			done
			return 0
		else
			return 4
		fi
	else
		return $RESULT
	fi
}

Script_Path=$(cd "$(dirname "$0")"; pwd)


if	[ ! -z $HELPPARAM ]
then
	help
	exit 2
else
	case $RUNMODE in
	O)
		wget_safe_onefile
		exit $?
	;;
	C)
		wget_safe_allfile
		exit $?
	;;
	P)
		wget_safe_allfile_phypath
		exit $?
	;;
	*)
		wget_safe
		exit $?
	;;
	esac
fi
