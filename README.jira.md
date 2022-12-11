# Setup Jira Docker

```shell
docker run --name="jira-legacy" -v /home/fuks-it/jira-app-data:/var/atlassian/application-data/jira -d -p 127.0.0.1:8080:8080 atlassian/jira-software:8.2.1
docker run --name="jira-legacy" -d -p 127.0.0.1:8080:8080 atlassian/jira-software:8.2.1

docker cp jira-app/atlassian-jira/WEB-INF/lib/mysql-connector-java-5.1.38-bin.jar 279c5e78e4d8:/opt/atlassian/jira/atlassian-jira/WEB-INF/lib
```

## Docker network

```shell
docker network create jira
docker network connect jira bf34f8b507e8
docker network inspect jira
```