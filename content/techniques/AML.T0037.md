---
atlas_id: AML.T0037
atlas_type: technique
attack_ref_id: T1005
attack_ref_url: https://attack.mitre.org/techniques/T1005/
created_date: "2021-05-13"
description: Злоумышленники могут искать в локальных источниках системы, таких как файловые системы, конфигурационные файлы или локальные базы данных, чтобы найти интересующие файлы и чувствительные данные перед эксфильтрацией....
generated: true
generated_by: atlasgen
maturity: realized
mitigation_count: 0
modified_date: "2026-05-27"
platforms:
    - Enterprise
procedure_count: 6
source_name: Data from Local System
subtechnique_count: 0
subtechnique_of: ""
tactics:
    - AML.TA0009
title: Данные из локальной системы
url: /techniques/AML.T0037/
---

Злоумышленники могут искать в локальных источниках системы, таких как файловые системы, конфигурационные файлы или локальные базы данных, чтобы найти интересующие файлы и чувствительные данные перед эксфильтрацией.

Это может включать базовые сведения для фингерпринтинга системы и чувствительные данные, например SSH-ключи.


## Тактики

<div class="relation-list">
<a class="relation-item" href="/tactics/AML.TA0009/"><span class="relation-id">AML.TA0009</span><strong>Сбор материалов</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0015/"><span class="relation-id">AML.CS0015</span><strong>Компрометация цепочки зависимостей PyTorch</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0009 Сбор материалов</span><p>Вредоносный пакет обследовал затронутую систему для базового фингерпринтинга, включая IP-адрес, имя пользователя и текущий рабочий каталог, а также похищал дополнительные чувствительные данные:</p><ul><li>DNS-серверы из `/etc/resolv.conf`</li><li>имя хоста из `gethostname()`</li><li>текущее имя пользователя из `getlogin()`</li><li>имя текущего рабочего каталога из `getcwd()`</li><li>переменные окружения</li><li>`/etc/hosts`</li><li>`/etc/passwd`</li><li>первые 1000 файлов в каталоге `$HOME`</li><li>`$HOME/.gitconfig`</li><li>`$HOME/.ssh/*`.</li></ul></a>
<a class="relation-item" href="/studies/AML.CS0043/"><span class="relation-id">AML.CS0043</span><strong>Прототип вредоносного ПО со встроенной промпт-инъекцией</strong><span class="relation-meta">Актор: Unknown Threat Actor / Тактика: AML.TA0009 Сбор материалов</span><p>Вредоносное ПО Skynet пытается собрать файлы `%HOMEPATH%\.ssh\known_hosts` и `C:/Windows/System32/Drivers/etc/hosts`.</p></a>
<a class="relation-item" href="/studies/AML.CS0044/"><span class="relation-id">AML.CS0044</span><strong>LAMEHUG: вредоносное ПО, использующее команды, динамически генерируемые ИИ</strong><span class="relation-meta">Актор: APT28 / Тактика: AML.TA0009 Сбор материалов</span><p>LAMEHUG использовал команды, сгенерированные ИИ, для сбора сведений о системе с сохранением в `%PROGRAMDATA%\info\info.txt`, а также рекурсивно просматривал папки Documents, Desktop и Downloads, чтобы подготовить файлы к эксфильтрации.</p></a>
<a class="relation-item" href="/studies/AML.CS0061/"><span class="relation-id">AML.CS0061</span><strong>AI in the Middle: веб-сервисы ИИ как ретрансляторы C2</strong><span class="relation-meta">Актор: Check Point Research / Тактика: AML.TA0009 Сбор материалов</span><p>Имплант собирал базовую информацию о хосте из локальной системы. Исследователи отметили, что это можно расширить до сбора таких сведений, как имя пользователя, домен, имя компьютера, установленное ПО, запущенные процессы и программы автозагрузки.</p></a>
<a class="relation-item" href="/studies/AML.CS0068/"><span class="relation-id">AML.CS0068</span><strong>Автономные агенты OpenAI, задействованные в оценочных испытаниях, скомпрометировали инфраструктуру Hugging Face</strong><span class="relation-meta">Актор: Autonomous OpenAI Agents / Тактика: AML.TA0009 Сбор материалов</span><p>Внешние ссылки HDF5 привели к раскрытию содержимого файла `/proc/self/environ` и файлов с исходным кодом воркера, в том числе сведений о том, как обрабатывались конфигурации наборов данных.</p></a>
<a class="relation-item" href="/studies/AML.CS0069/"><span class="relation-id">AML.CS0069</span><strong>GTG-1002 Claude Code Espionage Campaign</strong><span class="relation-meta">Актор: GTG-1002 / Тактика: AML.TA0009 Сбор материалов</span><p>Access obtained through the initial compromise also allowed GTG-1002&#39;s jailbroken Claude agent to gather credentials, system configurations, and sensitive operational data stored on compromised systems.</p></a>
</div>
