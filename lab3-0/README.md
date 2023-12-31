# LAB 3-0: EMR Cluster Creation

## Detalles del curso

| Información | |
| --- | --- |
| Nombre del curso | Tópicos especiales en Telemática |
| ID del curso | ST0263 |
| Nombre | Kevin Alejandro Sossa Chavarria (kasossac@eafit.edu.co) |
| Profesor | Edwin Nelson Montoya Munera (emontoya@eafit.edu.co) |

# 1. Objetivo principal

Creación de un clúster AWS EMR en Amazon para trabajar con todos los laboratorios.

---

# 2. Problemas resueltos y no resueltos

- [x] Creación de un bucket de AWS S3 para almacenar los notebooks que crearemos en el cluster de AWS EMR.
- [x] Creación de un cluster de AWS EMR.
- [x] Conexión SSH al master node.
- [x] Servicio Hue funcional.
- [x] Servicio JupyterHub funcional.
      
---

# 3. Entorno de ejecución

## Guía de uso

### Paso 1: Creación de un bucket AWS S3

Tendremos que crear un bucket de AWS S3 para almacenar los blocs de notas que crearemos en el clúster de AWS EMR.
    
1. En el servicio S3 haz clic en `Create bucket`:    
2. Introduce un nombre para el bucket, deja las opciones por defecto y dar click en `Create Bucket`.
   
---

### Paso 2: Creación del clúster de AWS EMR

1. En el servicio EMR haz clic en `Crear clúster`.
2. Introduzca un nombre para el clúster, seleccione la versión `emr-6.14.0` y en Application Bundle seleccione `Custom`. Despues selecciona las siguientes aplicaciones y activa `Use for Hive table metadata` y `Use for Spark table metadata`.
  
   ![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/b71db379-c657-4772-a47c-7affb7374b7f)
   
3. Edite la sección **Cluster termination** y asigne la finalización después de un tiempo de inactividad de tres horas.

   ![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/4a84e585-5ba8-4f2e-987d-ce098c77a213)

    
4. Deje las opciones siguiente como predeterminadas y vaya a la sección **Software configuration**.
5. Seleccione la opción `Enter configuration` y pegue la siguiente configuración:

   ```json
    [
      {
        "Clasificación": "jupyter-s3-conf",
        "Properties": {
          "s3.persistence.enabled": "true",
          "s3.persistence.bucket": "<NOMBRE_DE_TU_CUBO>"
        }
      }
    ]
    ```

6. Vaya a la sección **Security configuration and EC2 key pair**. Elige la vockey utilizada para los laboratorios (o podrías seleccionar una personalizada).
    
![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/a6b8628e-b8eb-437e-b9c6-a1baa928dcf4)

7. En la sección '**Identity and Access Management (IAM) roles** seleccione las siguientes opciones:

    - **Rol de Servicio de Amazon EMR**: EMR_DefaultRole
    - **Perfil de Instancia EC2 para Amazon EMR**: EMR_EC2_DefaultRole
    - **Rol de autoescalado personalizado**: LabRole

![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/4d4cc72d-be72-4eb2-a610-00571f2fb404)

8. Haga clic en `Crear cluster`.
---

### Paso 3: Editar grupos de seguridad para las aplicaciones

1. Dentro del menú de Amazon EMR selecciona la opción **Bloquear acceso público**.    
2. Selecciona el botón `Editar`. En la sección **Bloquear acceso público** seleccione la opción `Turn off` y seleccione el botón `Save`.
    
![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/51681fa9-e5b2-4c77-b75f-9d21bd5bc92e)

3. Dentro del menú de Amazon EMR ve a **Clusters** y selecciona el `Cluster ID` que tiene el estado `Esperando`. A continuación selecciona la opción `Applications`.
    
![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/9f53b8ed-826f-46b4-87b2-c482f67a7fe3)

5. Busca el servicio EC2. A continuación, vaya a la sección `Grupos de seguridad`.
    
6. Haz clic en el ID del SG que tiene el nombre `ElasticMapReduce-master`.
    
![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/02ec3abf-30bb-496c-9841-4e766be88121)

7. Haz clic en `Edit inbound rules`.
    
![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/f2ff3909-c48f-4e10-bb25-2da7f56f83bc)

8. Para cada uno de los puertos mostrados en los pasos 4, realice lo siguiente:
    1. Haga clic en `Add rule`:
    2. Selecciona la opción `Custom TCP`) e introduce el número de puerto:
    3. Haga clic en "Save rules" cuando haya añadido todas.

---

### Sección 4: Acceso al nodo primario

1. Dentro del menú de Amazon EMR vaya a **Clusters** y seleccione el `Cluster ID` que tiene el crear antes.
2. Haga clic en la URL `Connect to the Primary node using SSM` y siga las instrucciones que allí se indican. 
    
![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/a2701bd0-1da6-42c6-80cc-bee2ac18c81a)

3. Una conexión SSH exitosa al master node del clúster tendrá el siguiente aspecto:
    
![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/71d56bbc-e420-4124-b675-3bcbe536ed18)

4. Tendrá que editar el archivo 'hue.ini' siguiendo estos pasos: 
    1. Escribe el siguiente comando en el terminal:
   
        ```bash
        sudo vim /etc/hue/conf/hue.ini
        ```
        
    2. Busca la línea que contiene 'webhdfs_url' y cambia el puerto. Debes poner el puerto del Nodo de Nombre HDFS que se encuentra en la sección de Aplicaciones sobre tu cluster. 
    3. Guarda los cambios sobre el archivo.
    4. Reinicia el servicio Hue utilizando el siguiente comando:
        
        ```bash
        sudo systemctl restart hue.service
        ```
        
---

### Sección 5: Uso de Hue

1. En el menú de Amazon EMR, vaya a **Clusters** y seleccione el `Cluster ID` que tenga el estado `Waiting`. A continuación seleccione la opción `Applications`.    
2. Selecciona la URL del campo `Hado`. A continuación, introduce el usuario `hadoop` y una contraseña de tu elección.
    
![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/f23a0cb0-d86a-4a66-a54e-759174ed70a4)

---

### Parte 7: Uso de JupyterHub

1. Inicie sesión en la consola web de AWS y busque el servicio EMR, seleccione el `Cluster ID` que tiene el estado 'Waiting'; a continuación, seleccione la opción **Applications**.
2. Selecciona la URL del campo 'JupyterHub'. Ahora, introduce el usuario `jovyan` y la contraseña `jupyter`.
3. Selecciona el botón `New` y elige la opción `PySpark`:    
4. Verifica que las variables de contexto de Spark están instaladas:
    
![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/74b243bf-1d2e-43f6-8425-959ab794a389)
