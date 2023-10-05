**Komiser is back 🎉 and we're looking for maintainers to work on the new [roadmap](https://roadmap.tailwarden.com/), if you're interested, join us on our <a href="https://discord.tailwarden.com">Discord</a> community**

---

<h1 align="center"><img src="https://cdn.komiser.io/images/readme-komiser-header.png?version=latest" alt="Amp Logo"></h1>

Komiser is an open-source cloud-agnostic resource manager designed to analyze and manage cloud cost, usage, security, and governance all in one place. It integrates seamlessly with multiple cloud providers, including AWS, Azure, Civo, Digital Ocean, OCI, Linode, Tencent, Scaleway and [more](#supported-cloud-providers). Interested? read more about Komiser on our [website](https://komiser.io?utm_source=github&utm_medium=social).

<p align="center">
    
<a href="https://discord.tailwarden.com">
<img alt="Discord" src="https://img.shields.io/discord/932683789384183808?color=5865F2&label=Discord&style=for-the-badge" />
</a>

<a href="https://github.com/tailwarden/komiser/releases">
<img alt="GitHub commit activity" src="https://img.shields.io/github/commit-activity/m/tailwarden/komiser?color=00FF0&style=for-the-badge">
</a>

<a href="https://hub.docker.com/r/mlabouardy/komiser">
<img alt="Docker pulls" src="https://img.shields.io/docker/pulls/mlabouardy/komiser.svg?color=4591DF&style=for-the-badge">
</a>

</p>

<h4 align="center">
    <a href="https://cloud.tailwarden.com?utm_source=github&utm_medium=social">Tailwarden Cloud</a> |
    <a href="https://docs.komiser.io/docs/introduction/getting-started?utm_source=github&utm_medium=social/">Guide</a> |
    <a href="https://docs.komiser.io/docs/guides/overview?utm_source=github&utm_medium=social">How to Komiser</a> |
    <a href="https://docs.komiser.io/docs/events">Community Events</a><br/><br/>
    <a href="https://docs.komiser.io/docs/contributing/contribute?utm_source=github&utm_medium=social">Contribute</a> | 
    <a href="https://roadmap.tailwarden.com">Roadmap</a><br/><br/>
</h4>

<p align="center">
  <a href="https://www.youtube.com/watch?v=Hmerp_HMEf8">
    <img src=".github/images/deploy-to-aws.png" width="137" />
  </a>
</p>


<h1 align="center"><img src="https://cdn.komiser.io/gifs/readme-komiser-repo.gif?version=latest" alt="Komiser gif"></h1>

The fastest and most reliable way to get started with Komiser is signing up for free to [Tailwarden Cloud](https://cloud.tailwarden.com?utm_source=github&utm_medium=social)

[![Twitter URL](https://img.shields.io/twitter/url/https/twitter.com/fold_left.svg?style=social&label=Follow%20%40Komiser)](https://twitter.com/komiserdotio) [![Tweet](https://img.shields.io/twitter/url/http/shields.io.svg?style=social)](https://twitter.com/intent/tweet?text=Optimize%20Cost%20and%20Security%20on%20AWS&url=https://github.com/tailwarden/komiser&via=mlabouardy&hashtags=komiser,aws,gcp,cloud,serverless,devops) 

## Komiser CLI, try it out! 🚀
---
The easiest way to get started with Komiser is to install the [Komiser CLI](https://docs.komiser.io/docs/introduction/getting-started?utm_source=github&utm_medium=social) by running the `Homebrew` commands below. Don't have Homebrew? Install it [here](https://docs.brew.sh/Installation).

```
brew tap tailwarden/komiser
brew install komiser
```

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [What is Komiser? 🤷](#what-is-komiser-)
  - [Who is using it?](#who-is-using-it)
- [Getting started 👇](#getting-started-)
  - [Download](#download)
  - [Installation on AWS](#installation-on-aws)
    - [Connect Komiser CLI to your AWS account.](#connect-komiser-cli-to-your-aws-account)
    - [Deploy Komiser to single account access EKS cluster (Helm chart)](#deploy-komiser-to-single-account-access-eks-cluster-helm-chart)
    - [Deploy Komiser to a multi account access EKS cluster (Helm chart)](#deploy-komiser-to-a-multi-account-access-eks-cluster-helm-chart)
  - [Installation on Azure](#installation-on-azure)
  - [Installation on GCP](#installation-on-gcp)
  - [Installation on Civo](#installation-on-civo)
  - [Installation on OCI](#installation-on-oci)
  - [Installation on Digital Ocean](#installation-on-digital-ocean)
  - [Installation on Linode](#installation-on-linode)
  - [Installation on Tencent Cloud](#installation-on-tencent-cloud)
  - [Installation on Scaleway](#installation-on-scaleway)
- [Documentation 📖](#documentation-)
  - [Jump right in:](#jump-right-in)
- [Bugs and feature requests 🐞](#bugs-and-feature-requests-)
- [Roadmap and Contributing 🛣️](#roadmap-and-contributing-️)
- [Users 🧑‍🤝‍🧑](#users-)
- [Versioning 🧮](#versioning-)
- [Contributors](#contributors)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

## What is Komiser? 🤷
Komiser is an open source project created to **analyse** and **manage cloud cost**, **usage**, **security** and **governance** all in one place. With komiser you can also: 
* Build an inventory of your cloud infrastructure assets.
* Control your **resource usage** and gain visibility across all used services to achieve maximum cost-effectiveness.
* Detect **potential vulnerabilities** that could put your cloud environment at risk.
* Get a deep understanding of **how you spend** on the AWS, Azure, GCP, Civo, OVH, DigitalOcean and OCI.
* Uncover idle and untagged resources, ensuring that no resource goes unnoticed.

<h1 align="center"><img src=".github/images/dashboard.png" alt="Komiser dashboard"></h1>

## Who is using it?
Komiser was built with every Cloud Engineer, Developer, DevOps engineer and SRE in mind. We understand that tackling cost savings, security improvements and resource usage analyse efforts can be hard, sometimes just knowing where to start, can be the most challenging part at times. Komiser is here to help those cloud practitioners see their cloud resources and accounts much more clearly. Only with clear insight can timely and efficient actions take place.

## Supported cloud providers

## Getting started

### Installation

#### Linux

```
wget https://cli.komiser.io/latest/komiser_Linux_x86_64 -O komiser
```

#### Windows

```
wget https://cli.komiser.io/latest/komiser_Windows_x86_64.zip
tar -xf komiser_Windows_x86_64.zip
```

#### Mac OS X

ARM architecture (M1 & M2 Chip)
```
wget https://cli.komiser.io/latest/komiser_Darwin_arm64 -O komiser
```

AMD architecture (Intel Chip)
```
wget https://cli.komiser.io/latest/komiser_Darwin_x86_64 -O komiser
```

#### Homebrew
```
brew update
brew tap tailwarden/komiser
brew install komiser

# Make sure you are running the newest version of Komiser:
brew update
brew reinstall komiser
```

### Configuration

| Providers                                                                                                          | Documentation                                                                                                                                                    | Video |
| ------------------------------------------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---- |
| [![Kubernetes](https://img.shields.io/badge/kubernetes-%23326ce5.svg?style=for-the-badge&logo=kubernetes&logoColor=white)](https://docs.komiser.io/configuration/cloud-providers/k8s?utm_source=github&utm_medium=social&utm_campaign=readme)        | [Kubernetes documentation link](https://docs.komiser.io/configuration/cloud-providers/k8s?utm_source=github&utm_medium=social&utm_campaign=readme) | |
| [![AWS](https://img.shields.io/badge/AWS-%23FF9900.svg?style=for-the-badge&logo=amazon-aws&logoColor=white)](https://docs.komiser.io/configuration/cloud-providers/aws?utm_source=github&utm_medium=social&utm_campaign=readme)                      | [AWS documentation link ](https://docs.komiser.io/configuration/cloud-providers/aws?utm_source=github&utm_medium=social&utm_campaign=readme)           | |
| [![Azure](https://img.shields.io/badge/azure-%230078D4.svg?style=for-the-badge&logo=microsoftazure&logoColor=white)](https://docs.komiser.io/configuration/cloud-providers/azure?utm_source=github&utm_medium=social&utm_campaign=readme)        | [Azure documentation link](https://docs.komiser.io/configuration/cloud-providers/azure?utm_source=github&utm_medium=social&utm_campaign=readme) | |
| [![Civo](https://img.shields.io/badge/Civo-%23239DFF.svg?style=for-the-badge&logo=civo&logoColor=white)](https://docs.komiser.io/configuration/cloud-providers/civo?utm_source=github&utm_medium=social&utm_campaign=readme)                      | [Civo documentation link](https://docs.komiser.io/configuration/cloud-providers/civo?utm_source=github&utm_medium=social&utm_campaign=readme)           | |
| [![Digital Ocean](https://img.shields.io/badge/DigitalOcean-%230080FF.svg?style=for-the-badge&logo=digitalocean&logoColor=white)](https://docs.komiser.io/configuration/cloud-providers/digital-ocean?utm_source=github&utm_medium=social&utm_campaign=readme)        | [Digital Ocean documentation link](https://docs.komiser.io/configuration/cloud-providers/digital-ocean?utm_source=github&utm_medium=social&utm_campaign=readme) | |
| [![Akamai (Linode)](https://img.shields.io/badge/Akamai-%230096D6.svg?style=for-the-badge&logo=akamai&logoColor=white)](https://docs.komiser.io/configuration/cloud-providers/linode?utm_source=github&utm_medium=social&utm_campaign=readme)                      | [Akamai (Linode) documentation link](https://docs.komiser.io/configuration/cloud-providers/linode?utm_source=github&utm_medium=social&utm_campaign=readme)           | |
| [![GCP](https://img.shields.io/badge/GCP-%234285F4.svg?style=for-the-badge&logo=googlecloud&logoColor=white)](https://docs.komiser.io/docs/cloud-providers/google-cloud-platform?utm_source=github&utm_medium=social&utm_campaign=readme)                      | [GCP documentation link](https://docs.komiser.io/docs/cloud-providers/google-cloud-platform?utm_source=github&utm_medium=social&utm_campaign=readme)           | |
| [![OCI](https://img.shields.io/badge/OCI-%23F80000.svg?style=for-the-badge&logo=oracle&logoColor=white)](https://docs.komiser.io/configuration/cloud-providers/oracle?utm_source=github&utm_medium=social&utm_campaign=readme)        | [Oracle Cloud Infrastructure documentation link](https://docs.komiser.io/configuration/cloud-providers/oracle?utm_source=github&utm_medium=social&utm_campaign=readme) | |
| [![Tencent Cloud](https://img.shields.io/badge/Tencent-%230096D6.svg?style=for-the-badge&logo=tencentcloud&logoColor=white)](https://docs.komiser.io/configuration/cloud-providers/linode?utm_source=github&utm_medium=social&utm_campaign=readme)                      | [Tencent Cloud documentation link](https://docs.komiser.io/configuration/cloud-providers/linode?utm_source=github&utm_medium=social&utm_campaign=readme)           | |
| [![Scaleway](https://img.shields.io/badge/Scaleway-%234F0599.svg?style=for-the-badge&logo=scaleway&logoColor=white)](https://docs.komiser.io/docs/cloud-providers/scaleway?utm_source=github&utm_medium=social&utm_campaign=readme)        | [Scaleway documentation link](https://docs.komiser.io/docs/cloud-providers/scaleway?utm_source=github&utm_medium=social&utm_campaign=readme) | |
| [![MongoDB Atlas](https://img.shields.io/badge/MongoDB-%2347A248.svg?style=for-the-badge&logo=mongodb&logoColor=white)](https://docs.komiser.io/configuration/cloud-providers/linode?utm_source=github&utm_medium=social&utm_campaign=readme)                      | [MongoDB Atlas documentation link](https://docs.komiser.io/configuration/cloud-providers/linode?utm_source=github&utm_medium=social&utm_campaign=readme)           | |


## Resources 📖

* [Documentation](https://docs.komiser.io?utm_source=github&utm_medium=social)
* [Videos](https://www.youtube.com/@tailwarden/videos)


## Bugs and feature requests 🐞

Have a bug or a feature request? Please first read the issue guidelines and search for existing and closed issues. If your problem or idea is not addressed yet, [please open a new issue](https://github.com/tailwarden/komiser/issues).

## Roadmap and Contributing 🛣️

We are very excited about what is in store in the coming weeks and months, take a look at the [public roadmap](https://roadmap.tailwarden.com/) to stay on top of what's coming down the pipeline. 

Komiser is written in `Golang` and is `Elv2 licensed` - contributions are always welcome whether that means providing feedback through GitHub, through the `#feedback` channel on our [Discord server](https://discord.tailwarden.com), testing existing features or suggesting new ones. Feel free to check out our [contributor guidelines](./CONTRIBUTING.md) and consider becoming a **contributor** today. 

Learn how to contribute with these walkthrough videos:

- [How to contribute to Komiser engine](https://www.youtube.com/watch?v=Vn5uc2elcVg)
- [How to contribute to Komiser dashboard](https://www.youtube.com/watch?v=uwxj11-eRt8)

## Users 🧑‍🤝‍🧑

If you'd like to have your company represented and are using `Komiser` please give formal written permission below via email to contact@tailwarden.com.

We will need a URL to an SVG or png logo, a text title, and a company URL.

## Contributors

<a href="https://github.com/tailwarden/komiser/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=tailwarden/komiser" />
</a>

## Versioning 🧮

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/tailwarden/komiser/tags). 

## License

Komiser is licensed under the terms of [Elastic License 2.0 (ELv2)](https://github.com/tailwarden/komiser/blob/release/LICENSE).
