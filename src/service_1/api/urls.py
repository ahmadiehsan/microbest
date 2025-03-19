from django.urls import path

from api.views import API

urlpatterns = [path("", API.urls)]
