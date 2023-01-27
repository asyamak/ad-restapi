# Advertisement Api #

It's a service for storing and submitting ads.
For storing I used Postgres. Application runs over HTTP in JSON format.

To initialize application type in the console:

`make dcomposebuild`

then

`make dcomposeup`

There are 4 endpoints for ad manipulations.

### Create ad:
`/ad/create`

You must fill out request body like example below:
```
{
    "name": "ad name",
    "description": "ad description",
    "price": 3000,
    "photo_links":[
       "link1",
       "link2",
       "link3"
    ]
}
```
There are no more than 3 photo links, no more than 1000 symbols for description and no more than 200 symbols for ad's name.


### Get ads

`/ads`

Getting ads handler is implementing filter for ad's date of creation: `asc/desc`, ad's price: `asc/desc` which are must be used separately. Ads are displayed by 10 items per page, which also can be scrolled.
Example of body:
```
{
    "page": 1,
    "price": "desc",
    "date": ""
}
```
or
```
{
    "page": 1,
    "price": "",
    "date": "asc"
}
```

### Get ad
`/ad`

Getting ad handler displays particular ad by guid.

```
{
    "id":"20eb4229-9b06-416e-a293-b1a0501028bc"
}   
```

### Delete ad
`/as/delete`

Delete ad handler deletes ad by ad's guid.
```
{
    "id": "a19041eb-5490-4062-b320-43b1954c45d5"
}
```
