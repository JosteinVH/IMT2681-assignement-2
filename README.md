# Assignment #2 IMT2681 Cloud Technologies

### Description
Similar to Assignemt 1 - the in memory storage is moved to mongodb - and webhook implemented.

Submitted track files is and webhooks, as of requirement, are being stored in mongodb. Reason for this is because it's benefical instead of data being deleted e.g from a system crash, or reboot. This makes the application dependent on the database. Time is precious, and the cost it takes to query data from database is noticeable.
The performance overall is reduced, to the benefit of data availability.

**goigc** is used for processing of the igc data https://github.com/marni/goigc

### Heroku
xxx
### Openstack
xxx

---

## Testing

Tests can be found in mongodb folder.
Only testing the database functions.

---
## API Requirement
 #### GET /paragliding/
* Should redirect to /paragliding/ap

---

 #### GET /paragliding/api
* Returns meta information about the API 
* Body template:
```json
{
  "uptime": "<uptime>",
  "info": "Service for IGC tracks.",
  "version": "v1"
}
```
**\<uptime>** is the current uptime of the service formatted according to Duration format as specified by ISO 8601.


---

#### POST /paragliding/api/igc

* What: igc fil
* Request body template

```json
{
    "url" : "<link to igc file>"
}
```
 * Response body:
```json
{
    "id" : "<id of the igc file>"
}
```
* where: **\<url>** represents a normal URL, that would work in a browser, eg: http://skypolaris.org/wp-content/uploads/IGS%20Files/Madrid%20to%20Jerez.igc and **\<id>** represents an ID of the track, according to your internal management system.
---

#### GET /api/igc
* What: returns the array of all tracks ids
* Response: the array of IDs, or an empty array if no tracks have been stored yet.

```json
["<id1>", "<id2>", "..."]
```
---

#### GET /api/igc/\<id>

* What: returns the meta information about a given track with the provided \<id>, or NOT FOUND response code with an empty body.
* Response: 
```json
{
    "H_date": "<date from File Header, H-record>",
    "pilot": "<pilot>",
    "glider": "<glider>",
    "glider_id": "<glider_id>",
    "track_length": "<calculated total track length>"
}
```

---


#### GET /api/igc/\<id>/\<field>

* What: returns the single detailed meta information about a given track with the provided <id>, or NOT FOUND response code with an empty body.
* Response
 
    * **\<pilot>** for pilot
    * **\<glider>** for glider
    * **\<glider_id>** for glider_id
    * **\<calculated total track length>** for track_length
    * **\<H_date>** for H_date



---
HERHAHRHRHRR
#### GET /api/ticker/latest 

* What: returns the meta information about a given track with the provided \<id>, or NOT FOUND response code with an empty body.
* Response: 
  
* **\<timestamp>** for the latest added track 
---

---

#### GET /api/ticker/

* What: returns the JSON struct representing the ticker for the IGC tracks. The first track returned should be the oldest. The array of track ids returned should be capped at 5, to emulate "paging" of the responses.
* Response: 
```json
{
"t_latest": <latest added timestamp>,
"t_start": <the first timestamp of the added track>, this will be the oldest track recorded
"t_stop": <the last timestamp of the added track>, this might equal to t_latest if there are no more tracks left
"tracks": [<id1>, <id2>, ...],
"processing": <time in ms of how long it took to process the request>
}
```

---

---

#### GET /api/ticker/\<timestamp>

* What: What: returns the JSON struct representing the ticker for the IGC tracks. The first returned track should have the timestamp HIGHER than the one provided in the query. The array of track IDs returned should be capped at 5, to emulate "paging" of the responses.
* Response: 
```json
{
   "t_latest": <latest added timestamp of the entire collection>,
   "t_start": <the first timestamp of the added track>, this must be higher than the parameter provided in the query
   "t_stop": <the last timestamp of the added track>, this might equal to t_latest if there are no more tracks left
   "tracks": [<id1>, <id2>, ...],
   "processing": <time in ms of how long it took to process the request>
}
```

---

---

## Webhooks API

#### POST /api/webhook/new_track/

* What: Registration of new webhook for notifications about tracks being added to the system. Returns the details about the registration.
* Response: 
```json
{
    "webhookURL": {
      "type": "string"
    },
    "minTriggerValue": {
      "type": "number"
    }
}
```

---

#### GET /api/webhook/new_track/\<webhook_id>

* What: Accessing registered webhooks. Registered webhooks should be accessible using the GET method and the webhook id generated during registration.
* Response: 
```json
{
    "webhookURL": {
      "type": "string"
    },
    "minTriggerValue": {
      "type": "number"
    }
}
```

---


#### DELETE /api/webhook/new_track/\<webhook_id>

* What: What: Deleting registered webhooks. Registered webhooks can further be deleted using the DELETE method and the webhook id.
* Response: 
```json
{
    "webhookURL": {
      "type": "string"
    },
    "minTriggerValue": {
      "type": "number"
    }
}
```
---



#### Clock trigger

* What: Checks every 10min if the number of tracks differs from the previous check, and if it does, it will notify a predefined Slack webhook. 

---

#### Admin API 


#### GET /admin/api/tracks_count

* What:  Returns the current count of all tracks in the DB


---


#### DELETE /admin/api/tracks

* What: deletes all tracks in the DB
* Response: count of the DB records removed from DB
---










