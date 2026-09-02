---
atlas_id: AML.TA0009
atlas_type: tactic
attack_ref_id: TA0009
attack_ref_url: https://attack.mitre.org/tactics/TA0009/
created_date: "2022-01-24"
description: Злоумышленник пытается собрать ИИ-артефакты и другую связанную информацию, имеющую отношение к его цели. Сбор материалов включает техники, с помощью которых злоумышленники могут собирать информацию и выявлять...
generated: true
generated_by: atlasgen
modified_date: "2025-04-09"
procedure_count: 25
source_name: Collection
technique_count: 10
title: Сбор материалов
url: /tactics/AML.TA0009/
---

Злоумышленник пытается собрать ИИ-артефакты и другую связанную информацию, имеющую отношение к его цели.

Сбор материалов включает техники, с помощью которых злоумышленники могут собирать информацию и выявлять источники, из которых можно получить сведения, необходимые для дальнейшего достижения их целей.

После сбора данных следующей целью злоумышленников часто становится кража (эксфильтрация) ИИ-артефактов или использование собранной информации для подготовки будущих операций.

К распространённым целевым источникам относятся репозитории программного обеспечения, реестры контейнеров, репозитории моделей и объектные хранилища.


## Техники

<div class="relation-list">
<a class="relation-item" href="/techniques/AML.T0035/"><span class="relation-id">AML.T0035</span><strong>Сбор ИИ-артефактов</strong></a>
<a class="relation-item" href="/techniques/AML.T0036/"><span class="relation-id">AML.T0036</span><strong>Данные из информационных репозиториев</strong></a>
<a class="relation-item" href="/techniques/AML.T0037/"><span class="relation-id">AML.T0037</span><strong>Данные из локальной системы</strong></a>
<a class="relation-item" href="/techniques/AML.T0085/"><span class="relation-id">AML.T0085</span><strong>Данные из ИИ-сервисов</strong></a>
<a class="relation-item" href="/techniques/AML.T0085.000/"><span class="relation-id">AML.T0085.000</span><strong>Базы данных RAG</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0085.000/"><span class="relation-id">AML.T0085.000</span><strong>Базы данных RAG</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0085.001/"><span class="relation-id">AML.T0085.001</span><strong>Инструменты ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0085.001/"><span class="relation-id">AML.T0085.001</span><strong>Инструменты ИИ-агента</strong><span class="relation-meta">Подтехника</span></a>
<a class="relation-item" href="/techniques/AML.T0126/"><span class="relation-id">AML.T0126</span><strong>Automated Collection</strong></a>
<a class="relation-item" href="/techniques/AML.T0127/"><span class="relation-id">AML.T0127</span><strong>Data Staged</strong></a>
</div>


## Примеры процедур из кейсов

<div class="relation-list procedure-relations">
<a class="relation-item" href="/studies/AML.CS0006/"><span class="relation-id">AML.CS0006</span><strong>Ошибочная конфигурация Clearview AI</strong><span class="relation-meta">Актор: Researchers at spiderSilk / Тактика: AML.TA0009 Сбор материалов</span><p>Приватный репозиторий кода содержал учетные данные, которые использовались для доступа к облачным хранилищам AWS S3. Это привело к обнаружению ресурсов инструмента распознавания лиц, включая:</p><ul><li>выпущенные настольные и мобильные приложения</li><li>предварительные версии приложений с новыми возможностями</li><li>токены доступа Slack</li><li>необработанные видео и другие данные</li></ul></a>
<a class="relation-item" href="/studies/AML.CS0010/"><span class="relation-id">AML.CS0010</span><strong>Нарушение работы сервиса Microsoft Azure</strong><span class="relation-meta">Актор: Microsoft AI Red Team / Тактика: AML.TA0009 Сбор материалов</span><p>Команда нашла файл целевой ML-модели и необходимые обучающие данные.</p></a>
<a class="relation-item" href="/studies/AML.CS0015/"><span class="relation-id">AML.CS0015</span><strong>Компрометация цепочки зависимостей PyTorch</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0009 Сбор материалов</span><p>Вредоносный пакет обследовал затронутую систему для базового фингерпринтинга, включая IP-адрес, имя пользователя и текущий рабочий каталог, а также похищал дополнительные чувствительные данные:</p><ul><li>DNS-серверы из `/etc/resolv.conf`</li><li>имя хоста из `gethostname()`</li><li>текущее имя пользователя из `getlogin()`</li><li>имя текущего рабочего каталога из `getcwd()`</li><li>переменные окружения</li><li>`/etc/hosts`</li><li>`/etc/passwd`</li><li>первые 1000 файлов в каталоге `$HOME`</li><li>`$HOME/.gitconfig`</li><li>`$HOME/.ssh/*`.</li></ul></a>
<a class="relation-item" href="/studies/AML.CS0018/"><span class="relation-id">AML.CS0018</span><strong>Выполнение произвольного кода через Google Colab</strong><span class="relation-meta">Актор: Tony Piazza / Тактика: AML.TA0009 Сбор материалов</span><p>Злоумышленник может искать в системе жертвы частные и проприетарные данные, включая артефакты ML-моделей. Jupyter Notebook [позволяют выполнять shell-команды](https://colab.research.google.com/github/jakevdp/PythonDataScienceHandbook/blob/master/notebooks/01.05-IPython-And-Shell-Commands.ipynb). В этом примере смонтированный Drive проверяется на наличие checkpoint-файлов моделей PyTorch: &gt; /content/drive/MyDrive/models/checkpoint.pt</p></a>
<a class="relation-item" href="/studies/AML.CS0023/"><span class="relation-id">AML.CS0023</span><strong>ShadowRay: захват кластеров Ray, доступных из интернета</strong><span class="relation-meta">Актор: Unknown / Тактика: AML.TA0009 Сбор материалов</span><p>Злоумышленники могут собирать ИИ-артефакты, включая продакшен-модели и данные. Исследователи наблюдали рабочие продакшен-нагрузки нескольких организаций из разных отраслей.</p></a>
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0009 Сбор материалов</span><p>Промпт просит агента получить все поля и строки из «Customer Support Account Owners.csv». Агент извлекает весь файл.</p></a>
<a class="relation-item" href="/studies/AML.CS0037/"><span class="relation-id">AML.CS0037</span><strong>Эксфильтрация данных через инструменты ИИ-агента в Copilot Studio</strong><span class="relation-meta">Актор: Zenity / Тактика: AML.TA0009 Сбор материалов</span><p>Промпт просит агента получить все записи Salesforce с помощью инструмента `get-records`. Агент извлекает все записи из CRM организации-жертвы.</p></a>
<a class="relation-item" href="/studies/AML.CS0038/"><span class="relation-id">AML.CS0038</span><strong>Внедрение инструкций для отложенного автоматического вызова инструмента ИИ-агента</strong><span class="relation-meta">Актор: Embrace the Red / Тактика: AML.TA0009 Сбор материалов</span><p>Расширение Workspace находило документ и помещало его содержимое в контекст чата.</p></a>
<a class="relation-item" href="/studies/AML.CS0039/"><span class="relation-id">AML.CS0039</span><strong>Living Off AI: промпт-инъекция через Jira Service Management</strong><span class="relation-meta">Актор: Cato CTRL / Тактика: AML.TA0009 Сбор материалов</span><p>Вредоносный промпт предписывал собрать все сведения из других задач. Это вызывало инструмент Atlassian MCP, который мог обращаться к Jira-тикетам и собирать их.</p></a>
<a class="relation-item" href="/studies/AML.CS0043/"><span class="relation-id">AML.CS0043</span><strong>Прототип вредоносного ПО со встроенной промпт-инъекцией</strong><span class="relation-meta">Актор: Unknown Threat Actor / Тактика: AML.TA0009 Сбор материалов</span><p>Вредоносное ПО Skynet пытается собрать файлы `%HOMEPATH%\.ssh\known_hosts` и `C:/Windows/System32/Drivers/etc/hosts`.</p></a>
<a class="relation-item" href="/studies/AML.CS0044/"><span class="relation-id">AML.CS0044</span><strong>LAMEHUG: вредоносное ПО, использующее команды, динамически генерируемые ИИ</strong><span class="relation-meta">Актор: APT28 / Тактика: AML.TA0009 Сбор материалов</span><p>LAMEHUG использовал команды, сгенерированные ИИ, для сбора сведений о системе с сохранением в `%PROGRAMDATA%\info\info.txt`, а также рекурсивно просматривал папки Documents, Desktop и Downloads, чтобы подготовить файлы к эксфильтрации.</p></a>
<a class="relation-item" href="/studies/AML.CS0058/"><span class="relation-id">AML.CS0058</span><strong>Извлечение ИИ-моделей из Google Photos</strong><span class="relation-meta">Актор: Skyld / Тактика: AML.TA0009 Сбор материалов</span><p>Исследователи собрали артефакты моделей TensorFlow Lite из нескольких мест в APK, включая незашифрованные ресурсы приложения, файлы, встроенные в нативную библиотеку, и каталоги, специфичные для приложения.</p></a>
</div>


Показано 12 из 25 примеров.
