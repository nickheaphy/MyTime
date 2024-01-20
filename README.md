# MyTime

## What I want to capture

- Time/Duration
- Primary Log Type (ie Presales, Project, Postsales, Leave etc)
- Secondary Log Type (ie Presales-Solution Design, Presales-CustomerMeeting etc)
- Text Description
- Customer

When selecting the Primary Type, it should change the secondary log type to the appropriate fields. Ideally this would be radio buttons?

Primary and Secondary should be loaded from the database (though probably don't see a UI to actually enter, just write into the database)

## Installation (Mac)

This is a CGO free version of sqlite `go install modernc.org/sqlite@latest`

Need SQLite3 `brew install sqlite3`

### Ease of Use features

I was it to download the last X entries and put this into a dropdown(?) so rather than having to type "Solution Design for XXX" I can select and have it automatically populate all the same items as the last time I used that.

https://getbootstrap.com/docs/5.3/forms/form-control/#datalists


## Notes

https://earthly.dev/blog/golang-sqlite/
https://www.w3schools.com/xml/ajax_xmlhttprequest_send.asp
https://forum.golangbridge.org/t/using-golang-to-make-a-string-and-uses-ajax/22652/2
https://medium.com/@edwardpie/processing-form-request-data-in-golang-2dff4c2441be
