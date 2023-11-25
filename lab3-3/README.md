## Detalles del curso

| Informacion |  |
| --- | --- |
| Nombre del curso | Tópicos especiales en Telemática |
| ID del curso | ST0263 |
| Nombre | Kevin Alejandro Sossa Chavarria (kasossac@eafit.edu.co) |
| Profesor | Edwin Nelson Montoya Munera (emontoya@eafit.edu.co) |

# 1. Objetivo principal

Almacenar los datos del conjunto de datos del curso en un Datawarehouse y poder recuperar esos datos con consultas.

---

# 2. Problemas resueltos 

- [x] Creación de un clúster AWS redshift.
- [x] Ejecutar consultas básicas en la base de datos de demostración 'tickit'.
- [x] Crear un rol IAM para Amazon Redshift y ejecutar consultas para crear una base de datos externa
- [x] Creación de un cluster ERM y almacenar datos con hive

---


# 3. Entorno de ejecución

## Guía

Pasos a seguir para un correcto catálogo de los datos.

---

### Paso 1: Creación de un cluster de AWS redshift.

1. En el servicio de Redshift eelecciona la sección `Clusters` y haz clic en `Create cluster`.
2. Elige el nodo de tipo `dc2.large` y `1` número de nodos.

![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/8476a269-4f9e-4d23-a24c-fbdb29f9535b)

3. En la sección Datos de muestra pulse la casilla `Load sample data`.
4. En la configuración de la base de datos utilice el usuario `awsuser` y la contraseña `St1800232`.
5. En la sección de roles IAM asociados haga clic en `Manage IAM roles` y añada el rol `LabRole` y `myRedshiftRole`
6. Por ultimo haga click en `create cluster`

### Paso 2: Ejecutar consultas básicas en `tickit` database
1. Espere hasta que el cluster esté disponible y haga clic en el menú de lista `Query data` y haga clic en la opción `Query in query editor v2`.
2. Para conectarse a la base de datos, haga clic en la opción `Database username and password` y rellene los espacios en blanco con las credenciales indicadas anteriormente.

![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/f22df74f-039c-4b4a-a9e9-5de8cf9a18f5)

3. Entra en la carpeta `sample_data_dev` y pulsa el botón `tickit open samples queries` y pulsa crear
4. Ejecute las consultas de muestra

   ![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/185789cc-7b2d-46ce-9da4-91edb2cc949f)

## Paso 3: Crear un rol IAM para Amazon Redshift y ejecutar consultas para crear una base de datos externa en el bucket s3 de AWS

1. En el servicio `IAM Redshift` elija Roles y haga clic en `Crear rol`.
2. Elija `Redshift - Customizable` y haga clic en Siguiente.
3. Aparecerá la página Adjuntar política de permisos, añade `AmazonS3ReadOnlyAccess`, `AWSGlueConsoleFullAccess` y `AmazonAthenaFullAcces` y haz clic en `next`.
   
   ![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/f6006cc6-6fdf-4fdf-83bc-14782d4e6e89)

4. En Nombre de rol, introduzca `myspectrum_role`, y dar en `Create role`.
5. Crear la base de datos externa

![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/55f7644e-f688-46d7-b1f4-b80702ab014d)

6. Crear una tabla con datos externos en S3

   ![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/c058ca27-873c-4290-b59e-a9f65c8c233d)

7. Consultar datos

   ![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/6d564bb1-2477-43eb-9f82-19fc4a76a634)

## Paso 4: Creación de un cluster ERM y almacenamiento de datos con hive 

1. Crear un cluster ERM
2. Conéctate al nodo primario con ssh
3. Vamos a utilizar beeline que se utiliza para conectarse a Hive que se ejecuta en el clúster EMR y ejecutar HiveQL querie. Aquí está cómo se puede [conectar](https://sparkbyexamples.com/apache-hive/connect-to-hive-using-beeline/) a la misma
4. puedes ver todas las tablas que has creado

   ![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/1ed2c192-c28f-479e-847d-04ca7f08995e)

5. Crear tabla HDI en EMR/S3/Hue/Hive

    #tabla externa en S3:
    
    `CREATE EXTERNAL TABLE <nombre_de_tu_tabla> (ejemplo tus_datos:id INT) 
    FORMATO DE FILA CAMPOS DELIMITADOS TERMINADOS POR ',' 
    ALMACENADO COMO ARCHIVO DE TEXTO 
    LOCATION 's3://tu_s3bucket/'`

![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/550f0fe7-c940-40d9-9016-c6cce4de49d8)

6. Ahora puede consultar sus datos
   ![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/67bda9aa-ca17-44de-b4ac-483f1d4083c7)
