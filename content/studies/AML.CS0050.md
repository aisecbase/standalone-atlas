---
actor: DepthFirst
atlas_id: AML.CS0050
atlas_type: case-study
case_study_type: exercise
description: Исследователь безопасности продемонстрировал уязвимость удаленного выполнения кода (RCE) в один клик в ИИ-агенте OpenClaw через вредоносную ссылку, содержащую JavaScript-скрипт, выполнение которого занимает...
generated: true
generated_by: atlasgen
incident_date: "2026-02-01"
incident_date_granularity: Day
incident_date_raw: "2026-02-01"
procedure:
    - description: Исследователь разработал JavaScript-скрипт для RCE в один клик.
      description_line: Исследователь разработал JavaScript-скрипт для RCE в один клик.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0017
      technique_name: Разработка средств для атаки
    - description: Исследователь разместил вредоносный скрипт на неприметном сайте.
      description_line: Исследователь разместил вредоносный скрипт на неприметном сайте.
      tactic: AML.TA0003
      tactic_name: Подготовка ресурсов
      technique: AML.T0079
      technique_name: Размещение средств атаки
    - description: Когда жертва переходила по ссылке на сайт исследователя, вредоносный JavaScript-скрипт выполнялся в ее браузере.
      description_line: Когда жертва переходила по ссылке на сайт исследователя, вредоносный JavaScript-скрипт выполнялся в ее браузере.
      tactic: AML.TA0005
      tactic_name: Выполнение
      technique: AML.T0011.003
      technique_name: Вредоносная ссылка
    - description: Вредоносный скрипт открывал фоновое окно с интерфейсом управления OpenClaw жертвы, указывая в `gatewayUrl` WebSocket-адрес сервера исследователя. Интерфейс управления OpenClaw доверяет параметру `gatewayUrl` без проверки и автоматически подключается при загрузке, отправляя Gateway-токен на сервер исследователя.
      description_line: Вредоносный скрипт открывал фоновое окно с интерфейсом управления OpenClaw жертвы, указывая в `gatewayUrl` WebSocket-адрес сервера исследователя. Интерфейс управления OpenClaw доверяет параметру `gatewayUrl` без проверки и автоматически подключается при загрузке, отправляя Gateway-токен на сервер исследователя.
      tactic: AML.TA0013
      tactic_name: Доступ к учетным данным
      technique: AML.T0106
      technique_name: Эксплуатация уязвимостей для доступа к учетным данным
    - description: Вредоносный скрипт использовал Cross-Site WebSocket Hijacking (CSWSH), чтобы обойти сетевые ограничения localhost. Он открывал новое WebSocket-соединение с сервером OpenClaw Gateway на localhost.
      description_line: Вредоносный скрипт использовал Cross-Site WebSocket Hijacking (CSWSH), чтобы обойти сетевые ограничения localhost. Он открывал новое WebSocket-соединение с сервером OpenClaw Gateway на localhost.
      tactic: AML.TA0007
      tactic_name: Уклонение от защиты
      technique: AML.T0107
      technique_name: Эксплуатация уязвимостей для обхода защиты
    - description: Вредоносный скрипт использовал похищенный Gateway-токен для аутентификации, что позволяло затем выполнять вызовы к OpenClaw Gateway API в системе жертвы.
      description_line: Вредоносный скрипт использовал похищенный Gateway-токен для аутентификации, что позволяло затем выполнять вызовы к OpenClaw Gateway API в системе жертвы.
      tactic: AML.TA0012
      tactic_name: Повышение привилегий
      technique: AML.T0012
      technique_name: Действующие учетные записи
    - description: |-
        Вредоносный скрипт отключал защитную функцию OpenClaw, которая запрашивает подтверждение пользователя перед выполнением потенциально опасных команд. Для этого в OpenClaw Gateway API отправлялась следующая полезная нагрузка:

        ```json
        { "method": "exec.approvals.set",
          "params": { "defaults": { "security": "full", "ask": "off" } }
        }
        ```
      description_line: 'Вредоносный скрипт отключал защитную функцию OpenClaw, которая запрашивает подтверждение пользователя перед выполнением потенциально опасных команд. Для этого в OpenClaw Gateway API отправлялась следующая полезная нагрузка: ```json { "method": "exec.approvals.set", "params": { "defaults": { "security": "full", "ask": "off" } } } ```'
      tactic: AML.TA0007
      tactic_name: Уклонение от защиты
      technique: AML.T0081
      technique_name: Изменение конфигурации ИИ-агента
    - description: Вредоносный скрипт отключал песочницу OpenClaw, заставляя агента выполнять команды напрямую на хост-машине, а не внутри Docker-контейнера. Для этого в OpenClaw Gateway API отправлялся запрос `config.patch`, устанавливающий `tools.exec.host` в значение `"gateway"`.
      description_line: Вредоносный скрипт отключал песочницу OpenClaw, заставляя агента выполнять команды напрямую на хост-машине, а не внутри Docker-контейнера. Для этого в OpenClaw Gateway API отправлялся запрос `config.patch`, устанавливающий `tools.exec.host` в значение `"gateway"`.
      tactic: AML.TA0012
      tactic_name: Повышение привилегий
      technique: AML.T0105
      technique_name: Выход на хост
    - description: Вредоносный скрипт добивался удаленного выполнения кода, отправляя запрос `node.invoke` (RPC-механизм OpenClaw) к API OpenClaw.
      description_line: Вредоносный скрипт добивался удаленного выполнения кода, отправляя запрос `node.invoke` (RPC-механизм OpenClaw) к API OpenClaw.
      tactic: AML.TA0005
      tactic_name: Выполнение
      technique: AML.T0050
      technique_name: Интерпретатор команд и сценариев
procedure_count: 9
references:
    - title: 1-Click RCE To Steal Your Moltbot Data and Keys (CVE-2026-25253)
      url: https://depthfirst.com/post/1-click-rce-to-steal-your-moltbot-data-and-keys
    - title: CVE-2026-25253
      url: https://nvd.nist.gov/vuln/detail/CVE-2026-25253
    - title: openclaw
      url: https://openclaw.ai/blog/introducing-openclaw
reporter: ""
source_name: OpenClaw 1-Click Remote Code Execution
target: OpenClaw
title: Удаленное выполнение кода (RCE) в OpenClaw в один клик
url: /studies/AML.CS0050/
---

Исследователь безопасности продемонстрировал уязвимость удаленного выполнения кода (RCE) в один клик в ИИ-агенте OpenClaw через вредоносную ссылку, содержащую JavaScript-скрипт, выполнение которого занимает миллисекунды. Об этой уязвимости было сообщено; для затронутых версий OpenClaw она отслеживается как `CVE-2026-25253`.[1] OpenClaw описывается как персональный ИИ-ассистент, запускаемый на устройствах пользователя: он отвечает в уже используемых чат-приложениях, а в отличие от SaaS-ассистентов, где данные находятся на чужих серверах, OpenClaw запускается там, где выбирает пользователь: на ноутбуке, в домашней лаборатории или на VPS. «Ваша инфраструктура. Ваши ключи. Ваши данные».[2]

Исследователь показал, что когда жертва переходит по вредоносной ссылке, в ее браузере выполняется клиентский JavaScript-скрипт, который может украсть токены аутентификации из интерфейса управления OpenClaw через WebSocket-соединение. Затем скрипт использует Cross-Site WebSocket Hijacking, чтобы обойти ограничения localhost для OpenClaw Gateway API. После установления соединения он аутентифицируется с украденным токеном и изменяет конфигурацию агента OpenClaw: отключает подтверждение пользователя и обеспечивает выход за пределы контейнера, что позволяет запускать команды оболочки напрямую на хост-машине.

[1]: https://nvd.nist.gov/vuln/detail/CVE-2026-25253
[2]: https://openclaw.ai/blog/introducing-openclaw
