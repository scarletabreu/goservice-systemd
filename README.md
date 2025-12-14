# Go Service con systemd

## Descripción
Este proyecto consiste en el desarrollo y despliegue de un servicio escrito en Go, administrado mediante systemd en una máquina virtual Linux. El objetivo es aplicar conceptos fundamentales de administración de sistemas operativos, gestión de procesos, seguridad y automatización de servicios en entornos Linux.

## Tecnologías utilizadas
- Golang (Go)
- Linux (Ubuntu)
- systemd
- SSH
- tmux

## Usuarios del sistema
- Usuario administrador: `scarlet`
- Usuario de servicio: `goservice`

El servicio se ejecuta bajo un usuario sin privilegios administrativos, mientras que las tareas de configuración y administración se realizan desde el usuario administrador.

## Seguridad
- Acceso remoto mediante SSH usando autenticación por clave pública.
- Inicio de sesión del usuario root deshabilitado.
- Autenticación por contraseña deshabilitada para SSH.
- Separación de privilegios entre usuario administrador y usuario de servicio.

## Compilación del proyecto
Desde el directorio del proyecto:

```bash
go build -o goservice
```

## Configuración del servicio systemd
El servicio se define mediante un archivo goservice.service, ubicado en `/etc/systemd/system/`, con reinicio automático ante fallos. Los comandos principales:

```bash
sudo systemctl daemon-reload
sudo systemctl start goservice
sudo systemctl enable goservice
sudo systemctl status goservice
```

## Monitoreo y gestión
Para la visualización de logs:

```bash
journalctl -u goservice
```

Para la administración de sesiones:

```bash
tmux
```

Para el monitoreo de procesos:

```bash
htop
```

## Estructura del proyecto

```bash
goservice-systemd/
├── cmd/
│   └── goservice/
│       └── main.go
├── systemd/
│   └── goservice.service
├── docs/
│   └── screenshots/
├── go.mod
├── go.sum
├── .gitignore
└── README.md
```

## Autor
Scarlet Abreu
