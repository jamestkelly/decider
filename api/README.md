# Decider | REST API Documentation

<a name="readme-top"></a>

<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/jamestkelly/decider/api">
    <img src="../resources/img/decider_logo_background.png" alt="Logo" width="80" height="80">
  </a>

<h3 align="center">Decider</h3>

  <p align="center">
    A progressive web application (PWA) for coming to a consensus on decisions, complaints and suggestions sharing, and creating events.
    <br />
    <a href="https://github.com/jamestkelly/decider"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/jamestkelly/decider">View Demo</a>
    ·
    <a href="https://github.com/jamestkelly/decider/issues">Report Bug</a>
    ·
    <a href="https://github.com/jamestkelly/decider/issues">Request Feature</a>
  </p>
</div>

# Overview

This `README` serves the purpose of providing a static list of REST APIs as implemented for the `Decider` application.

## REST API

The REST APIs utilised by the Decider application are described below for the base URL, `/api`.

### Register User

#### Request

`POST /register`

    curl -i -H 'Accept: application/json' -d 'email=example@gmail.com&password=something' 'http://localhost:5000/register'

#### Response

    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    Content-Length: 2

    []

## Validate User Token

### Request

`POST /validateToken`

    curl -i -H 'Accept: application/json' -d 'token=231274613497' http://localhost:5000/validateToken

### Response

    HTTP/1.1 201 Created
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 201 Created
    Connection: close
    Content-Type: application/json
    Location: /thing/1
    Content-Length: 36

    {"id":231274613497,"email":"example@gmail.com&password=something"}

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/jamestkelly/decider.svg?style=for-the-badge

[contributors-url]: https://github.com/jamestkelly/decider/graphs/contributors

[forks-shield]: https://img.shields.io/github/forks/jamestkelly/decider.svg?style=for-the-badge

[forks-url]: https://github.com/jamestkelly/decider/network/members

[stars-shield]: https://img.shields.io/github/stars/jamestkelly/decider.svg?style=for-the-badge

[stars-url]: https://github.com/jamestkelly/decider/stargazers

[issues-shield]: https://img.shields.io/github/issues/jamestkelly/decider.svg?style=for-the-badge

[issues-url]: https://github.com/jamestkelly/decider/issues

[license-shield]: https://img.shields.io/github/license/jamestkelly/decider.svg?style=for-the-badge

[license-url]: https://github.com/jamestkelly/decider/blob/master/LICENSE.txt