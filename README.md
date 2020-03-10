# Useragent

Useragent formatter and parser based on User-Agent guidelines at Kiwi.com.

## About

It is difficult to debug API requests (such as malformed requests or overloading)
when the source of these requests is not easy to identify;
user agents should make identification of the source much faster and easier.

At Kiwi.com it is mandatory for every HTTP request made to an internal API
to include a *User-Agent* header in following format:

`<service name>/<version> (Kiwi.com <environment>) [<system info>]`

- where `<service name>` is name of the service making the request.
- where `<version>` is either the git commit hash (preferred), or a version number which is tagged in git.
- where `<environment>` should be a string that matches the environment reported to Datadog, Datadog APM, and Sentry.
In most cases that would mean production.
It can differ only when different environment strings are sent to these services for filtering hacks or similar reasons.
- where `[<system info>]` may be set to provide the standard info specified in user agent strings,
such as requests library, language, and OS versions.

Examples:

- `User-Agent: rambo/87f2594 (Kiwi.com production)`
- `User-Agent: donut-service/b5bf54c32a1bb27aa45543a9eac0affb8d8b32a5 (Kiwi.com Joe-Doe-testing) requests/3.3.3 python/2.8.0`

Validation of the user agent is recommended to be done with the following regex:

- `^(?P<name>.+?)\/(?P<version>.+?) \(Kiwi\.com (?P<environment>.+?)\)(?: ?(?P<system_info>.*))$`

Or this JavaScript variant: 

- `/^(.+[^\/])\/(.+?) \(Kiwi\.com (.+?)\)(?: ?(?:.*))$/`

You can try this pattern in an interactive sandbox with the above examples at [https://regex101.com/r/Jcaw67/3](https://regex101.com/r/Jcaw67/3)

## Go-Useragent

Go-useragent module provides few utility functions for parsing and formatting the _User-Agent_ strings
in previously defined format. This allows you to for example tag metrics only with the service name or
separate logs by individual environments.
