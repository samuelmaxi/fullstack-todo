from rest_framework import serializers
from .models import Task

class TaskSerializer(serializers.Serializer):
  id = serializers.IntegerField(read_only=True)
  name = serializers.CharField(max_length=200)
  isDone = serializers.BooleanField(default=False, source='is_done')
  
  def create(self, validated_data):
    return Task.objects.create(**validated_data)
  
  def update(self, instance, validated_data):
    instance.name = validated_data.get('name', instance.name)
    instance.is_done = validated_data.get('is_done', instance.is_done)
    instance.save()
    return instance
    

class TaskDeSerializer(serializers.Serializer):
  name = serializers.CharField(max_length=200)
  is_done = serializers.BooleanField(default=False)
  
  def create(self, validated_data):
    return Task.objects.create(**validated_data)
  
  def update(self, instance, validated_data):
    instance.name = validated_data.get('name', instance.name)
    instance.is_done = validated_data.get('is_done', instance.is_done)
    instance.save()
    return instance