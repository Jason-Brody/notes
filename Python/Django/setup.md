### Django

+ 安装
> pip3 install django

+ 新建Django工程
> django-admin startproject \<project-name>

+ 启动Django工程
> python3 manage.py runserver

+ 创建app
    1. 新建 
    > python3 manage.py startapp \<app-name>
    2. 注册app(settings.py)
    ``` python
    INSTALLED_APPS = [
        '<app-name>', # 注册app
        'django.contrib.admin',
        'django.contrib.auth',
        'django.contrib.contenttypes',
        'django.contrib.sessions',
        'django.contrib.messages',
        'django.contrib.staticfiles',
    ]
    ```
    3. 全局url配置(urls.py中配置app url)
    ``` python
    from django.contrib import admin
    from django.urls import path,include

    urlpatterns = [
        path('admin/', admin.site.urls),
        path('<app-name>/',include('<app-name>.urls'))
    ]
    ```