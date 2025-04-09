from api.views import API
from django.urls import path

urlpatterns = [path("", API.urls)]
