---
atlas_id: AML.TA0014
atlas_type: tactic
attack_ref_id: TA0011
attack_ref_url: https://attack.mitre.org/tactics/TA0011/
created_date: "2024-04-11"
description: Злоумышленник пытается взаимодействовать со скомпрометированными ИИ-системами, чтобы управлять ими. Командование и управление включает техники, которые злоумышленники могут использовать для взаимодействия с системами...
generated: true
generated_by: atlasgen
modified_date: "2024-04-11"
procedure_count: 7
source_name: Command and Control
technique_count: 4
title: Командование и управление
url: /tactics/AML.TA0014/
---

Злоумышленник пытается взаимодействовать со скомпрометированными ИИ-системами, чтобы управлять ими.

Командование и управление включает техники, которые злоумышленники могут использовать для взаимодействия с системами под их контролем в сети организации-жертвы.
Злоумышленники обычно пытаются имитировать нормальный ожидаемый трафик, чтобы избежать обнаружения.
Существует множество способов, с помощью которых злоумышленник может установить канал командования и управления с разным уровнем скрытности в зависимости от структуры сети организации-жертвы и ее средств защиты.


## Техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0072/"><span class="relation-id">AML.T0072</span><strong>Реверс-шелл</strong></a>
<a class="relation-item" href="/techniques/AML.T0096/"><span class="relation-id">AML.T0096</span><strong>API ИИ-сервиса</strong></a>
<a class="relation-item" href="/techniques/AML.T0108/"><span class="relation-id">AML.T0108</span><strong>ИИ-агент</strong></a>
<a class="relation-item" href="/techniques/AML.T0114/"><span class="relation-id">AML.T0114</span><strong>Веб-интерфейс ИИ-сервиса</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0014 Командование и управление</span><p>Имплант Sliver предоставляет исследователю канал командного управления, позволяя изучать среду жертвы и продолжать атаку.</p></a>
<a class="relation-item" href="/studies/AML.CS0031/"><span class="relation-id">AML.CS0031</span><strong>Вредоносные модели на Hugging Face</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0014 Командование и управление</span><p>Вредоносная нагрузка представляла собой reverse shell, настроенный на подключение к жестко заданному IP-адресу.</p></a>
<a class="relation-item" href="/studies/AML.CS0042/"><span class="relation-id">AML.CS0042</span><strong>SesameOp: новый бэкдор использует OpenAI Assistants API как C2-канал</strong><span class="relation-meta">Актор: Unknown Threat Actor / Тактика: AML.TA0014 Командование и управление</span><p>Злоумышленник использовал OpenAI Assistants API как канал передачи команд вредоносному ПО SesameOp. SesameOp выполнял эти команды в системе жертвы и отправлял результаты обратно злоумышленнику по тому же каналу. Команды и результаты передавались в зашифрованном виде. Чтобы скрыть следы, SesameOp удалял объекты Assistants и Messages, которые создавал и использовал для связи.</p></a>
<a class="relation-item" href="/studies/AML.CS0051/"><span class="relation-id">AML.CS0051</span><strong>Использование OpenClaw для командования и управления через промпт-инъекцию</strong><span class="relation-meta">Актор: HiddenLayer / Тактика: AML.TA0014 Командование и управление</span><p>Промпт заставил OpenClaw действовать как C2-агент в интересах исследователя. OpenClaw запросил TODO-список с `https://openclaw.aisystem.tech/todo` с помощью своего навыка `web_fetch` и выполнил команды через свой навык `bash`.</p></a>
<a class="relation-item" href="/studies/AML.CS0052/"><span class="relation-id">AML.CS0052</span><strong>LLMSmith: уязвимости RCE в приложениях с интеграцией LLM</strong><span class="relation-meta">Актор: Researchers at University of Chinese Academy of Sciences, Shandong University, and University of New South Wales / Тактика: AML.TA0014 Командование и управление</span><p>Python-код открывал reverse shell, который использовался как канал командования и управления.</p></a>
<a class="relation-item" href="/studies/AML.CS0061/"><span class="relation-id">AML.CS0061</span><strong>AI in the Middle: веб-сервисы ИИ как ретрансляторы C2</strong><span class="relation-meta">Актор: Check Point Research / Тактика: AML.TA0014 Командование и управление</span><p>Имплант инициировал анонимную веб-сессию с публичным ИИ-сервисом и отправлял подготовленный промпт. Так формировался канал командования и управления, в котором данные эксфильтрировались через запросы к домену, подконтрольному злоумышленнику, а команды передавались обратно через ответ.</p></a>
<a class="relation-item" href="/studies/AML.CS0065/"><span class="relation-id">AML.CS0065</span><strong>Атака на цепочку поставок через повторное использование пространства имён модели</strong><span class="relation-meta">Актор: Unit 42 Researchers / Тактика: AML.TA0014 Командование и управление</span><p>Полезная нагрузка запустила на развёрнутом эндпоинте реверс-шелл, который установил соединение с подконтрольной исследователям инфраструктурой.</p></a>
</div>
