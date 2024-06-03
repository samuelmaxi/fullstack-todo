from django.db import models

# Create your models here.

class Task(models.Model):
  name = models.CharField(
    null=False,
    blank=False,
    max_length=250
  )
  
  def __str__(self) -> str:
    return self.name