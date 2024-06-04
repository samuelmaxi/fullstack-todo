from django.db import models

# Create your models here.

class Task(models.Model):
  name = models.CharField(
    null=False,
    blank=False,
    max_length=250
  )
  
  is_done = models.BooleanField(
    default=False,
  )
  
  def __str__(self) -> str:
    return f"{self.name} - {self.is_done}"