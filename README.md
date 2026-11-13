Server Manager Design Doc

This is a manager of devices and services. It should let me start, stop, add, remove services from devices that the server managers.

URL List

/home
This is the page that the user will interact with most. 
It will have:
    1. a list of devices by name and IP Address
    2. a list of services the device has added
    3. the ability to restart a service
    4. the ability to stop a service
    5. the ability to start a service
    6. the ability to get a service status
    7. a button to update all device and service statuses
It will need:
    1. An endpoint that calls the client to get a list of addable services
    2. An endpoint that calls the client to restart a service
    3. An endpoint that calls a client to start a service
    4. An endpoint that calls a client to stop a service
    5. An endpoint that calls a client to check service status

/device/ceate
This is the form for addding a device to the page. This works by adding the device to the database.
It will have:
    1. A form for submitting a device's IP and Name
    2. Redirect user to the home screen on successful form

/device/service/add
This is where a service is added to the device. 
It will have:
    1. a form for submitting the device name and service name
    2. redirect user to home screen on successful form
    3. If the device does not have the service, the form fails
It will need:
    1. an API endpoint that calls the device to see if it has the service name listed
    2. the endpoint should return if it came back successful or not

/service/create
This is a form that takes a device name and a service name, communicates with the device, and adds the service file to the server
It will have:



