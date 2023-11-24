# LAB 3-1: Gestión de archivos en HDFS y S3

## Detalles del curso

| Informacion |  |
| --- | --- |
| Nombre del curso | Tópicos especiales en Telemática |
| ID del curso | ST0263 |
| Nombre | Kevin Alejandro Sossa Chavarria (kasossac@eafit.edu.co) |
| Profesor | Edwin Nelson Montoya Munera (emontoya@eafit.edu.co) |

# 1. Main goal

Gestionar todos los archivos disponibles de los conjuntos de datos en los servicios HDFS y S3 ).

---

# 2. Problemas resueltos y no resueltos

- [x] Creación de un bucket de AWS S3 para almacenar los notebooks que crearemos en el cluster de AWS EMR.
- [x] Creación de un cluster de AWS EMR versión 6.13.0
- [x] Conexión SSH al nodo maestro.
- [x] Servicio Hue funcional.
- [x] Servicio JupyterHub funcional.

---
# 3. Entorno de ejecución

## Guía de uso
---

### Paso 1: Creación de un bucket AWS S3

Tendremos que crear un bucket de AWS S3 para almacenar los blocs de notas que crearemos en el clúster de AWS EMR.
  
1. En el servicio S3 haz clic en `Create bucket`:

2. Introduzca un nombre para el bucket, a continuación, en la sección `Object Ownership` seleccione la opción `ACLs enabled` y marque la casilla `Object writer`.

   A continuación, en la sección `Block Public Access settings for this bucket`, quita la marca de la opción `Block all public access` y marca la casilla `I acknowledge that the current settings might result in this bucket and the objects within becoming public.`:

![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/90b5a424-994f-4e60-b04f-9919c5e164dc)

![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/f03501e3-3452-44ac-a0e7-14735d71ef0c)

3. Deje las siguientes opciones por defecto y seleccione el botón `Create Bucket`:

4. en el menú S3, selecciona el nombre del bucket creado anteriormente, selecciona la sección `Permissions` y busca la subsección `Access control list (ACL))`, una vez allí, selecciona el botón `Edit`.

5. Marca las casillas `List` y `Read` para `Everyone (public access)` y `Authenticated users group (anyone with an AWS account)` luego marca la opción `I understand the effects of these changes on my objects and buckets` finalmente selecciona el botón `Save changes`:

![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/0c4d8c01-17c9-4ae8-8c85-a29ea244d8dd)

6. Vaya al [archivo de ejemplo](https://github.com/st0263eafit/st0263-232/blob/main/bigdata/datasets/airlines.csv) y seleccione el botón `Download raw file`:
7. Vuelva a la interfaz principal de buckets, seleccione el nombre del bucket creado anteriormente y arrastre el archivo descargado a él, después seleccione el botón `Upload`.

![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/f25c3c40-c2c3-44f1-9fe1-063395130e09)


8. Seleccione el nombre del archivo cargado anteriormente y, a continuación, en la sección `Propeties` de `Object overview` copie la `Object URL`.
9. En una ventana del navegador pegue la URL copiada anteriormente dejando a un lado '/airlines.csv'.

### Sección 2: Gestión de archivos en HDFS usando terminal

1. Tendremos que crear un cluster AWS EMR.
2. Una vez creado el clúster especificado anteriormente deberemos conectarnos a él vía SSH.
3. Una vez establecida la conexión con el nodo primario crearemos una carpeta llamada 'gutenberg-small' dentro de la ruta '/user/hadoop/datasets' utilizando los siguientes comandos:
    1. Crea el directorio 'datasets' dentro de la ruta 'user/hadoop/'.
   
    ```bash
    hdfs dfs -mkdir /usuario/hadoop/datasets
    ```

    2. Crea el directorio 'gutenber-small' dentro de la ruta 'user/hadoop/datasets/'.
   
    ```bash
    hdfs dfs -mkdir /usuario/hadoop/datasets/gutenberg-small
    ```

    3. Listar directorios y ficheros presentes en la ruta /user/hadoop/
   
    ```bash
    hdfs dfs -ls /usuario/hadoop/datasets
    ```
4. Pon el contenido de un local en el directorio '/user/hadoop/datasets/gutenberg-small' usando el siguiente comando. Para ello puedes clonar el repositorio que contiene los conjuntos de datos de ejemplo
    
    ```bash
    hdfs dfs -put <TU_CARPETA_LOCAL> /user/hadoop/datasets/gutenberg-small/
    ```
    
### Sección 3: Gestión de archivos en HDFS usando HUE

1. Vaya a la consola de AWS y busque el servicio EMR.
2. Selecciona el `Cluster ID` que tiene el estado `Waiting`; luego selecciona `options`.
3. Selecciona la URL del campo `Hue` e introduce el usuario `hadoop` y una contraseña.
4. Selecciona la sección `Archivos`:
5. Si ha seguido correctamente los pasos anteriores, debe haber encontrado la carpeta `datasets`. En caso contrario, crea la carpeta en el botón `New`:

![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/56b25243-7725-46a2-8898-d2aeaa2f047f)


6. Seleccione el botón `Upload`:

   ![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/95db96b0-53ec-4ff9-a29c-6c8767f5238f)

### Sección 4: Gestión de archivos en S3 usando HUE

1. Inicie sesión en la consola de AWS y busque el servicio EMR.
2. Seleccione el `Cluster ID` que tiene el estado`'Waiting`; luego seleccione `options`.
3. Selecciona la URL del campo `HUE` e introduce el usuario `hadoop` y una contraseña de tu elección.
4. Seleccione la sección `S3`:
5. Todos tus buckets deben encontrarse ahí.
6. Seleccione un bucket. A continuación, haga clic en el botón `Upload`, `Select Files` y, por último, seleccione los archivos que desea subir.
7. Verifique que el archivo es accesible a través de otras formas como AWS CLI. Ejecuta el siguiente comando para listar el archivo en tu bucket:

   ```bash
    aws s3 ls s3://<NOMBRE_Su_CAMBITO>
    ```

    Debería mostrar el archivo que subiste antes.
