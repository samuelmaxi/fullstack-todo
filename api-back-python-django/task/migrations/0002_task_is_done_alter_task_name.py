# Generated by Django 5.0.6 on 2024-06-04 22:15

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('task', '0001_initial'),
    ]

    operations = [
        migrations.AddField(
            model_name='task',
            name='is_done',
            field=models.BooleanField(default=False),
        ),
        migrations.AlterField(
            model_name='task',
            name='name',
            field=models.CharField(max_length=250),
        ),
    ]
