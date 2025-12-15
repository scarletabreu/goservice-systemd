# Desarrollo y Despliegue de un Servicio en Go con systemd

## Descripción
Este proyecto documenta el desarrollo, despliegue y administración de un servicio implementado en Go, ejecutado y gestionado mediante systemd en un entorno Linux (Ubuntu).  
El objetivo principal es aplicar y demostrar conceptos de administración de sistemas, gestión de procesos, seguridad y automatización de servicios en un entorno simulado de producción.

---

## Tecnologías y Componentes
- **Golang (Go):** Lenguaje de programación utilizado para el desarrollo del servicio TCP y de la aplicación secundaria `infoapp`.
- **Linux (Ubuntu LTS):** Sistema operativo de la máquina virtual empleada para el despliegue.
- **systemd:** Sistema de gestión de servicios y procesos.
- **UFW (Uncomplicated Firewall):** Herramienta utilizada para la configuración y gestión de políticas de firewall.
- **SSH:** Protocolo para acceso remoto seguro.
- **dpkg/deb:** Utilizado para empaquetar y distribuir la aplicación secundaria (`infoapp`).

---

## Usuarios del Sistema
- **Usuario administrador:** `scarlet` (responsable de la configuración y administración del servicio).  
- **Usuario de servicio sin privilegios:** `goservice` (bajo el cual se ejecuta el servicio Go para garantizar la separación de privilegios).  

El servicio principal se ejecuta con permisos restringidos, mientras que las tareas administrativas se realizan desde la cuenta del administrador.

---

## Configuración de Seguridad y Firewall
Las medidas de seguridad implementadas incluyen:

1. **Separación de privilegios:** El servicio `goservice` se ejecuta bajo un usuario específico sin privilegios administrativos.  
2. **Acceso SSH seguro:**  
   - Inicio de sesión del usuario `root` deshabilitado.  
   - Autenticación mediante clave pública; la autenticación por contraseña está deshabilitada.  
3. **Firewall activo (UFW):**  
   - Solo se permiten los puertos **22/tcp** (SSH) y **9090/tcp** (servicio Go).  
   - Todo el tráfico entrante no autorizado se bloquea por defecto.

---

## Aplicación Empaquetada `infoapp`
Se desarrolló una aplicación complementaria (`infoapp`) cuya función es mostrar la información de identificación del autor. La aplicación fue empaquetada en formato `.deb` para demostrar la distribución automatizada de software.

| Metadato | Valor |
| :--- | :--- |
| **Paquete** | `infoapp` |
| **Versión** | 1.0 |
| **Ubicación final** | `/usr/local/bin/infoapp` |

### Verificación de la Aplicación
Para comprobar la correcta instalación y funcionamiento:

```bash
infoapp
# Resultado esperado: ID: 10153953, Nombre: Scarlet Abreu, Correo: SXAS0002@ce.pucmm.edu.do
```

---

## Configuración y Administración del Servicio systemd
El servicio principal se define mediante el archivo `goservice.service`.

### Comandos de Gestión del Servicio
Para recargar la configuración, iniciar y habilitar el servicio al arranque:

```bash
sudo systemctl daemon-reload
sudo systemctl start goservice
sudo systemctl enable goservice
sudo systemctl status goservice
```

### Verificación de Persistencia y Disponibilidad
- El servicio utiliza la directiva `Restart=always` para garantizar su disponibilidad continua.  
- **Prueba de recuperación:** Tras forzar la terminación del proceso (`sudo systemctl kill --signal=SIGKILL goservice`), el servicio se reinicia automáticamente.  
- **Monitoreo de recursos:** Se aplican límites de CPU y memoria mediante CGroup, visibles mediante `htop`.

### 6.3 Monitoreo y Registro de Eventos
```bash
# Visualización de logs del servicio
journalctl -u goservice

# Gestión de sesiones persistentes
tmux

# Monitoreo de procesos y recursos del sistema
htop
```

---

## Estructura del Proyecto
La organización del repositorio final separa el código fuente de los artefactos de despliegue:

```bash
goservice-systemd/
├── cmd/
│   ├── goservice/                     # Código fuente del servicio TCP
│   │   └── main.go
│   └── infoapp/                       # Código fuente de la aplicación empaquetada
│       └── main.go
├── deb-packaging/
│   ├── infoapp-deb/                   # Estructura del paquete DEB
│   │   └── DEBIAN/
│   │       └── control
│   └── infoapp-deb.deb                # Paquete binario final
├── docs/
│   └── goservice-systemd.pdf          # Evidencia de pruebas
├── systemd/
│   └── goservice.service              # Unidad de servicio systemd
├── go.mod
├── .gitignore
└── README.md
```

---

## Autor
**Scarlet Abreu Sánchez**
