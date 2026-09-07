---
atlas_id: AML.T0025
atlas_type: technique
attack_ref_id: ""
attack_ref_url: ""
created_date: "2021-05-13"
description: Злоумышленники могут эксфильтровать ИИ-артефакты или другую информацию, важную для их целей, традиционными киберсредствами. Подробнее см. тактику ATT&CK Exfiltration.
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 1
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 14
source_name: Exfiltration via Cyber Means
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0010
title: Эксфильтрация киберсредствами
url: /techniques/AML.T0025/
---

Злоумышленники могут эксфильтровать ИИ-артефакты или другую информацию, важную для их целей, традиционными киберсредствами.

Подробнее см. тактику ATT&CK [Exfiltration](https://attack.mitre.org/tactics/TA0010/).


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0010/"><span class="relation-id">AML.TA0010</span><strong>Эксфильтрация</strong></a>
</div>


## Меры защиты

<div class="relation-list">
<a class="relation-item" href="/mitigations/AML.M0005/"><span class="relation-id">AML.M0005</span><strong>Контроль доступа к ИИ-моделям и хранимым данным</strong><p>Контроль доступа может предотвращать эксфильтрацию.</p></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0010/"><span class="relation-id">AML.CS0010</span><strong>Нарушение работы сервиса Microsoft Azure</strong><span class="relation-meta">Актор: Microsoft AI Red Team / Тактика: AML.TA0010 Эксфильтрация</span><p>Команда эксфильтровала модель и данные обычными киберсредствами.</p></a>
<a class="relation-item" href="/studies/AML.CS0015/"><span class="relation-id">AML.CS0015</span><strong>Компрометация цепочки зависимостей PyTorch</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0010 Эксфильтрация</span><p>Вся собранная информация, включая содержимое файлов, отправлялась через зашифрованные DNS-запросы на домен `*[dot]h4ck[dot]cfd` с использованием DNS-сервера `wheezy[dot]io`.</p></a>
<a class="relation-item" href="/studies/AML.CS0018/"><span class="relation-id">AML.CS0018</span><strong>Выполнение произвольного кода через Google Colab</strong><span class="relation-meta">Актор: Tony Piazza / Тактика: AML.TA0010 Эксфильтрация</span><p>Получив доступ к Google Drive, злоумышленник может открыть сервер для эксфильтрации частных данных или артефактов ML-моделей. В примере из исходной статьи показаны загрузка, установка и использование `ngrok`, серверного приложения, чтобы открыть доступный злоумышленнику URL к Google Drive жертвы и всем его файлам.</p></a>
<a class="relation-item" href="/studies/AML.CS0023/"><span class="relation-id">AML.CS0023</span><strong>ShadowRay: захват кластеров Ray, доступных из интернета</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0010 Эксфильтрация</span><p>ИИ-артефакты, учетные данные и другая ценная информация могут быть эксфильтрованы киберсредствами. Исследователи обнаружили признаки reverse shell на уязвимых кластерах; такие оболочки могут использоваться для закрепления, продолжения выполнения произвольного кода и эксфильтрации данных.</p></a>
<a class="relation-item" href="/studies/AML.CS0027/"><span class="relation-id">AML.CS0027</span><strong>Путаница с организациями на Hugging Face</strong><span class="relation-meta">Актор: threlfall_hax / Тактика: AML.TA0010 Эксфильтрация</span><p>Обнаруженные учетные данные могли быть эксфильтрованы через имплант Sliver.</p></a>
<a class="relation-item" href="/studies/AML.CS0043/"><span class="relation-id">AML.CS0043</span><strong>Прототип вредоносного ПО со встроенной промпт-инъекцией</strong><span class="relation-meta">Актор: Unknown Threat Actor / Тактика: AML.TA0010 Эксфильтрация</span><p>Вредоносное ПО Skynet настраивает Tor-прокси для эксфильтрации собранных файлов. Примечание: собранные файлы только выводились в stdout и фактически не были успешно эксфильтрованы.</p></a>
<a class="relation-item" href="/studies/AML.CS0044/"><span class="relation-id">AML.CS0044</span><strong>LAMEHUG: вредоносное ПО, использующее команды, динамически генерируемые ИИ</strong><span class="relation-meta">Актор: APT28 / Тактика: AML.TA0010 Эксфильтрация</span><p>LAMEHUG эксфильтровал собранные данные на серверы под контролем злоумышленников через SFTP или HTTP POST-запросы.</p></a>
<a class="relation-item" href="/studies/AML.CS0048/"><span class="relation-id">AML.CS0048</span><strong>Публично доступные интерфейсы управления ClawdBot позволили получить учётные данные и выполнить команды</strong><span class="relation-meta">Актор: Jamieson O&#39;Reilly / Тактика: AML.TA0010 Эксфильтрация</span><p>Исследователь мог бы использовать обнаруженные токены приложений, чтобы эксфильтрировать полные истории приватных переписок, включая переданные файлы, из любых подключенных мессенджеров: Telegram, Slack, Discord, Signal, WhatsApp и других.</p></a>
<a class="relation-item" href="/studies/AML.CS0058/"><span class="relation-id">AML.CS0058</span><strong>Извлечение ИИ-моделей из Google Photos</strong><span class="relation-meta">Актор: Skyld / Тактика: AML.TA0010 Эксфильтрация</span><p>Исследователи использовали статический анализ и инструментацию Frida для восстановления файлов моделей. Для зашифрованных моделей они перехватывали расшифрованные файлы TFLite во время выполнения, когда Google Photos загружал их для выполнения.</p></a>
<a class="relation-item" href="/studies/AML.CS0059/"><span class="relation-id">AML.CS0059</span><strong>EchoLeak: промпт-инъекция нулевого клика против M365 Copilot для эксфильтрации данных</strong><span class="relation-meta">Актор: Aim Labs / Тактика: AML.TA0010 Эксфильтрация</span><p>Чтобы обойти ограничения CSP, исследователи направили запрос к отображенному изображению через разрешенный путь предпросмотра или прокси Microsoft Teams, который загружал контролируемый злоумышленником URL с закодированным секретом.</p></a>
<a class="relation-item" href="/studies/AML.CS0063/"><span class="relation-id">AML.CS0063</span><strong>Атаки на Gemini с помощью промптов в приглашениях Google Calendar</strong><span class="relation-meta">Актор: SafeBreach Research Team / Тактика: AML.TA0010 Эксфильтрация</span><p>Подконтрольный злоумышленнику веб-сайт получил IP-адрес источника — устройства жертвы, что позволило приблизительно определить местоположение устройства.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0010 Эксфильтрация</span><p>Агенты передавали учётные данные, вывод команд, данные об окружении, отдельные строки приватных наборов данных и приватные архивы, связанные с заданиями, используя в качестве каналов репозитории наборов данных и плацдарм.</p></a>
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>Кибершпионская кампания GTG-1002 с использованием Claude Code</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0010 Эксфильтрация</span><p>After reviewing the summary, GTG-1002 approved the transfer of selected data over the Claude web service.</p></a>
<a class="relation-item" href="/studies/AML.CS0071/"><span class="relation-id">AML.CS0071</span><strong>Multi-Agent Framework Compromises Taiwanese Government Systems</strong><span class="relation-meta">Актор: Unknown Chinese-language actor / Тактика: AML.TA0010 Эксфильтрация</span><p>Over 2,564 personnel records, a complete user database, and internal architecture details were exfiltrated.</p></a>
</div>
