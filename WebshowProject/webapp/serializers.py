from rest_framework import serializers
from webapp.models import Project

class ProjectSerializer(serializers.ModelSerializer):
    class Meta:
        model = Project
        fields = ('projectid',
                'projectname',
                'description',
                'verified',
                'course',
                'projectimage',
                'projectvideo',
                'projectthumbnail',
                # 'userid'
                )