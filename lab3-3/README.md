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
