# daily

Very simple 50 line cron daemon for one job only

 * synchronously - only one job at one time
 * only daily mode - between 0-23h
 * if job exits with error the daemon will try again until it will success
 * all output to stdout/error
 
##### Motivation:

The Cron utility has a hard limitation, the daemon must to be runned under root-user. This is unacceptable for running docker containers under a non-root user.