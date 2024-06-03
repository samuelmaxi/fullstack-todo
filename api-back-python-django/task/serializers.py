from rest_framework import serializers
from .models import Task

class TaskSerializer(serializers.Serializer):
  id = serializers.IntegerField(read_only=True)
  name = serializers.CharField(max_length=200)
  
  def create(self, validated_data):
    return Task.objects.create(**validated_data)
  
  def update(self, instance, validated_data):
    instance.name = validated_data.get('name', instance.name)
    
class TaskDeSerializer(serializers.Serializer):
  name = serializers.CharField(max_length=200)
  
  def create(self, validated_data):
    return Task.objects.create(**validated_data)
  
  def update(self, instance, validated_data):
    instance.name = validated_data.get('name', instance.name)