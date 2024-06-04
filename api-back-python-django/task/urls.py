from django.urls import path
from . import views

app_name = "task"
urlpatterns = [
  path('task/', views.TaskListViewSet.as_view(), name='get'),
  path('task/search/<int:id>/', views.TaskMethodsAPIView.as_view(), name='get' ),
  path('task/create/', views.TaskMethodsAPIView.as_view(), name='post'),
  path('task/change/<int:id>/', views.TaskMethodsAPIView.as_view(), name='put'),
  path('task/edit/<int:id>/', views.TaskMethodsAPIView.as_view(), name='patch'),
  path('task/delete/<int:id>/', views.TaskMethodsAPIView.as_view(), name='delete'),
]