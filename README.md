## Laboratorio 2 - Sistemas Distribuidos

#### Integrantes

- Alfredo Llanos - 201804536-5
- Fernanda Cerda - 201804567-5
- Sofía Mañana - 201804535-7

#### Indicaciones

- El repositorio de todos los archivos en su última versión están en: https://github.com/nandink/Lab2_SD
##### Los programas están distribuidos de la siguiente manera:
 - dist29: se ejecuta el Lider y DataNode0. clave = <hZK%8XqG-?Naj!>~>
 - dist30: se ejecutan los jugadores y DataNode1. clave = <M*jvU4W;$#DZs_5,>
 - dist31: se ejecuta el NameNode. clave = <A3h!)7czKgsu{?GC>
 - dist32: se ejecuta el Pozo y DataNode2. clave = <a7+VT<F^*8mLvjX]>
##### Orden de ejecución:
 - En primer lugar, ejecutar el Lider (dentro de la carpeta Servicio) en la máquina dist29.
 - Luego, ejecutar el NameNode, DataNode's y el Pozo, cada uno en su respectiva máquina. Para ejecutar los DataNode's, se necesita abrir dos terminales por cada máquina donde estos corren.
 - Por último, ejecutar los Jugadores (dentro de la carpeta Cliente).
- Se encuentra implementado el primer juego (Red Light-Green Light). El programa termina cuando acaba dicho juego. 
- Las máquinas se van a encontrar listas para su ejecución, con la configuración del archivo .proto listo.

#### Consideraciones

- Tratamos de implementar la mayoría de la tarea pero no alcanzamos a probar, por lo que las funciones no probadas se 
encuentran comentadas en los lugares que les corresponden. La funcionalidad de estas si fue probada pero de manera local, no en las máquinas, por lo que las conexiones no están 100% confirmadas de que funcionan. Las dejamos comentadas para que la tarea en sí, logre compilar y no tengamos el 0 directo. Esperamos que sirva de algo.