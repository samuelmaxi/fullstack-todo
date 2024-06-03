from django.urls import path
from . import views

app_name = "task"
urlpatterns = [
  path('task/', views.TaskListViewSet.as_view(), name='get'),
  path('task/search/<int:id>/', views.TaskMethodsAPIView.as_view(), name='get' ),
  path('task/create/', views.TaskMethodsAPIView.as_view(), name='post'),
  # path('search/<int:>/'),
]