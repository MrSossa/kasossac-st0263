
## Detalles del curso

| Informacion |  |
| --- | --- |
| Nombre del curso | Tópicos especiales en Telemática |
| ID del curso | ST0263 |
| Nombre | Kevin Alejandro Sossa Chavarria (kasossac@eafit.edu.co) |
| Profesor | Edwin Nelson Montoya Munera (emontoya@eafit.edu.co) |

# 1. Objetivo principal

Catalogar los datos del conjunto de datos del curso y poder enviar consultas sin servidor

---

# 2. Problemas resueltos y no resueltos

- [x] Creación de un bucket S3 de AWS para almacenar los datos.
- [x] Creación de un crawler desde Glue para catalogar el conjunto de datos.
- [x] Consulta desde Athena sobre las tablas.
- [x] Almacenar los resultados de la consulta en S3.

---

# 3. Entorno de ejecución

## Guía

---

1. En el servicio de S3 selecciona el botón `Crear bucket`.
2. Introduce un nombre para el bucket; luego en la sección `Object Ownership` selecciona la opción `ACLs enabled` y marca la casilla `Object writer`. 
    
    A continuación, en la sección `Block Public Access settings for this bucket` quite la marca de la opción `Block all public access` y marque la casilla `‘I acknowledge that the current settings might result in this bucket and the objects within becoming public`:

3. Deje las siguientes opciones por defecto y seleccione el botón `Create Bucket`.
4. Descarga los [datasets](https://github.com/st0263eafit/st0263-232/tree/main/bigdata/datasets).
6. Vuelva a la interfaz principal de buckets, seleccione el nombre del bucket creado anteriormente y arrastre los datasets de datos descargados hasta él, luego dar click en `upload`.
  ![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/ccc57eb6-3684-4176-bd7d-b6cfd01c1539)


---

### Sección 2: Catalogar datos con AWS Glue

1. Vaya a la consola web de AWS y busque el servicio Glue
2. En la barra lateral de control, vaya a la sección `Crawlers`.
3. Haga clic en `Create crawler`.
4. Establezca un nombre para su crawler y haga clic en `Next`.
5. En la sección `Elegir fuentes de datos y clasificadores`, haga clic en `Add a data source`.
6. Navegue hasta encontrar la carpeta que desea catalogar.

   ![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/1f3ee3db-4f34-49c1-a89a-2246f3986416)

7. Una vez seleccionado, haga clic en `Add an S3 data source`.
8. La vista previa de la fuente de datos debería tener este aspecto. A continuación, haga clic en `Siguiente
9. En la sección `Configure security settings`, configure el LabRole.
10. Establezca la base de datos de destino por defecto. Mantenga la programación `On demand` para catalogar los datos manualmente.
11. Después de los pasos anteriores, el crawler debería crearse con éxito.
12. Ahora debe empezar a catalogar. Seleccione su crawler y haga click en `Run` para iniciar el job.
13. Una vez finalizado el trabajo, en la barra lateral de control vaya a la sección `Tables`.
14. La nueva tabla debe estar allí, con una ubicación de su cubo S3.

---

### Sección 3: Consulta con Athena

1. Vaya a la consola web de AWS y busque el servicio Athena.
2. En la barra lateral de control, seleccione la sección `Query editor`.
3. Debemos establecer una ubicación para almacenar nuestros resultados. Vaya a `Settings`. A continuación, haga clic en `Manage`.
4. En la `Location of query result` establezca una ruta sobre su bucket donde desea almacenar.
5. A continuación, haga clic en `Save`. La vista general de la configuración debería tener este aspecto.
6. Vuelva a la sección `Editor`. Seleccione la tabla creada con su crawler.

   ![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/9147cc2d-9157-4f11-b1ae-aa6b16e63953)

**Ejemplos
- Ejemplo de consulta:
```sql
  SELECT * FROM "default"."<TU_TABLE>" WHERE lifeex > 80 ORDER BY lifeex desc;
```
![imagen](https://github.com/MrSossa/kasossac-st0263/assets/83780739/7fbc8432-a8f7-43e5-988e-14a76b613072)

---

### Sección 4: Guardar los resultados de las consultas

1. En el servicio S3 seleccione su bucket.
2. Busque en la ruta que seleccionó en la Sección 3, paso 4. Esta ruta se define para almacenar la salida de sus consultas.
