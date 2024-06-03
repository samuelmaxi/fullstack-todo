from django.shortcuts import render
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status, viewsets, generics
from .models import Task
from .serializers import TaskSerializer
from django.core.exceptions import ObjectDoesNotExist

# Create your views here.

class TaskListViewSet(generics.ListAPIView):
  
  queryset = Task.objects.all()
  serializer_class = TaskSerializer
  
class TaskMethodsAPIView(APIView):
  
  def get(self, request, id=None):
    try:
      queryset = Task.objects.get(id=id)
      serializer = TaskSerializer(queryset)
      return Response(serializer.data, status=status.HTTP_200_OK)
    except Task.DoesNotExist:
      return Response(serializer.data, status=status.HTTP_404_NOT_FOUND)

  def post (self, request):
      serializer = TaskSerializer(data=request.data)
      if serializer.is_valid():
        serializer.save()
        return Response(serializer.data, status=status.HTTP_201_CREATED)
      return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)
    
