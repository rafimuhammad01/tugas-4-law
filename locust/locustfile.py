import time
from locust import HttpUser, task, between

npm = "1906398603"

class QuickstartUser(HttpUser):
    wait_time = between(1, 2.5)

    @task
    def read_only(self):
        self.client.get("/read/1906398603")

    @task
    def read_with_transaction(self):
        self.client.get("/read/1906398603/1655031730")
