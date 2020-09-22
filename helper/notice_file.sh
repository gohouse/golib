#!/bin/env sh
title="日常提醒"
content="记得喝水活动一下"
subtitle="记得喝水"
sound="Pon"
cmd=$(printf 'display notification "%s" with title "%s" subtitle "%s" sound name "%s"' "$content" "$title" "$subtitle" "$sound")
osascript -e "$cmd"
say -v Ting-ting $content