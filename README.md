# Sublime Cancel

This repository hosts a minimal reproduction of an issue I'm seeing on macOS where canceling a build in Sublime Text does not appear to send a kill signal when cancelling a build, which prevents the possibility of graceful shutdown.

It looks like Sublime _should_ be sending a SIGTERM if we look at [exec.py](https://github.com/twolfson/sublime-files/blob/master/Packages/Default/exec.py#L122) in the default package, so I'm not sure what's happening.

Thanks to keith and OdatNurd for helping me debug this far.

---

If I open the sublime project `repro.sublime-project`, set the `Run` build command as the build system, then press Cmd+B followed my Ctrl+C to cancel, I see the following output in Sublime's build results:

```
Process started. PID: 7805
Send signals to test (e.g., kill -SIGTERM 7805 )
Press Ctrl+C to send SIGINT
Press Ctrl+\ to send SIGQUIT

[Cancelled]
```

If I `cd` to the directory then run `go run .` followed by Ctrl-C, I see the following (note the `Caught signal`): 

```
Process started. PID: 7872
Send signals to test (e.g., kill -SIGTERM 7872 )
Press Ctrl+C to send SIGINT
Press Ctrl+\ to send SIGQUIT
^C
Caught signal: interrupt
exit status 130
```

I see the same result if I launch Sublime Text in safe mode (`subl --safe-mode`).