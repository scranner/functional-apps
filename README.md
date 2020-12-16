# Predictive Verification

This repo contains the code and applications required to test and verify the pythia mechanism.

Please view the guides in the [mono-repo wiki](https://gitlab.eeecs.qub.ac.uk/40182178/final-year-mono/wikis/home) for instruction in how to use this repo.

This repo is broken into the following folders:

## Grafana
Contains the dashboards and deployment configuration for a grafana instance, to allow visualisation of the pythia apps performance.
To install use the following command:
```
helm install grafana stable/grafana -f grafana.gcp.values.yaml
```
To login use the credentials `username: admin` `password: astrongpassword`:

You must then connect to the prometheus server, to do this select add datasource and input the prometheus instance name. If the default configuration has been used the url `http://pythia-prometheus-server` can be used for the datasource url.

To add the dashboard, from the right, hover over the plus mark and select import from the menu. Paste the `dashboard.json` here to import the dashboard.

## Load tests

Contains gatling tests that can be run using sbt. To run the tests, from the folder run the following command:
```
sbt "gatling:testOnly simulations.LoadTest"
sbt "gatling:testOnly simulations.ScalingTest"
```
This will output a report of the run scenario in the `/target/gatling` folder.
## Results
Final data and results from different tests run used for the final year report.
## Under Test
Applications that can be deployed to a kubernetes cluster with the mechanism enabled or disabled to allow testing of the mechanism.
## Verification
### scripts
Contains javascript scripts to download test data for use with statistical verification. To use, update prometheus url in main.js, and run the following commands
```
npm ci
node main.js
```

This will download the test data from the prometheus instance for the supplied pods and parameters.
### statistical_verification
Dockerized tests to allow the testing of multiple models. do build use
```
docker build . -t ver-tests
```

To run, you must pass the dataset name you wish to use and the number of timeseries to test. Run the following:

```
docker run -e dataset_name=2mWindow_2sStep -e dataset_size=1000 ver-tests
```