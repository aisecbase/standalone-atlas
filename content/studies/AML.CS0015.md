---
actor: Unknown
atlas_id: AML.CS0015
atlas_type: case-study
case_study_type: ""
description: Пакеты Linux для предварительной версии PyTorch под названием Pytorch-nightly были скомпрометированы с 25 по 30 декабря 2022 года из-за вредоносного двоичного файла, загруженного в репозиторий кода Python Package...
generated: true
generated_by: atlasgen
incident_date: ""
incident_date_granularity: ""
incident_date_raw: ""
procedure:
    - description: |-
        Вредоносный пакет зависимости с именем `torchtriton` был загружен в репозиторий PyPI с тем же именем, что и пакет, поставлявшийся со сборкой PyTorch-nightly.

        Этот вредоносный пакет содержал дополнительный код, отправлявший чувствительные данные с машины.

        Вредоносный `torchtriton` устанавливался вместо легитимного пакета, потому что PyPI имел приоритет над другими источниками. Подробнее см. [этот GitHub issue](https://github.com/pypa/pip/issues/8606).
      description_line: Вредоносный пакет зависимости с именем `torchtriton` был загружен в репозиторий PyPI с тем же именем, что и пакет, поставлявшийся со сборкой PyTorch-nightly. Этот вредоносный пакет содержал дополнительный код, отправлявший чувствительные данные с машины. Вредоносный `torchtriton` устанавливался вместо легитимного пакета, потому что PyPI имел приоритет над другими источниками. Подробнее см. [этот GitHub issue](https://github.com/pypa/pip/issues/8606).
      tactic: AML.TA0004
      tactic_name: Первичный доступ
      technique: AML.T0010.001
      technique_name: ПО для ИИ
    - description: |-
        Вредоносный пакет обследовал затронутую систему для базового фингерпринтинга, включая IP-адрес, имя пользователя и текущий рабочий каталог, а также похищал дополнительные чувствительные данные:

        - DNS-серверы из `/etc/resolv.conf`

        - имя хоста из `gethostname()`

        - текущее имя пользователя из `getlogin()`

        - имя текущего рабочего каталога из `getcwd()`

        - переменные окружения

        - `/etc/hosts`

        - `/etc/passwd`

        - первые 1000 файлов в каталоге `$HOME`

        - `$HOME/.gitconfig`

        - `$HOME/.ssh/*`.
      description_line: 'Вредоносный пакет обследовал затронутую систему для базового фингерпринтинга, включая IP-адрес, имя пользователя и текущий рабочий каталог, а также похищал дополнительные чувствительные данные: - DNS-серверы из `/etc/resolv.conf` - имя хоста из `gethostname()` - текущее имя пользователя из `getlogin()` - имя текущего рабочего каталога из `getcwd()` - переменные окружения - `/etc/hosts` - `/etc/passwd` - первые 1000 файлов в каталоге `$HOME` - `$HOME/.gitconfig` - `$HOME/.ssh/*`.'
      tactic: AML.TA0009
      tactic_name: Сбор материалов
      technique: AML.T0037
      technique_name: Данные из локальной системы
    - description: Вся собранная информация, включая содержимое файлов, отправлялась через зашифрованные DNS-запросы на домен `*[dot]h4ck[dot]cfd` с использованием DNS-сервера `wheezy[dot]io`.
      description_line: Вся собранная информация, включая содержимое файлов, отправлялась через зашифрованные DNS-запросы на домен `*[dot]h4ck[dot]cfd` с использованием DNS-сервера `wheezy[dot]io`.
      tactic: AML.TA0010
      tactic_name: Эксфильтрация
      technique: AML.T0025
      technique_name: Эксфильтрация киберсредствами
procedure_count: 3
references:
    - title: PyTorch statement on compromised dependency
      url: https://pytorch.org/blog/compromised-nightly-dependency/
    - title: Analysis by BleepingComputer
      url: https://www.bleepingcomputer.com/news/security/pytorch-discloses-malicious-dependency-chain-compromise-over-holidays/
reporter: PyTorch
source_name: Compromised PyTorch Dependency Chain
target: PyTorch
title: Компрометация цепочки зависимостей PyTorch
url: /studies/AML.CS0015/
---

Пакеты Linux для предварительной версии PyTorch под названием Pytorch-nightly были скомпрометированы с 25 по 30 декабря 2022 года из-за вредоносного двоичного файла, загруженного в репозиторий кода Python Package Index (PyPI). Вредоносный двоичный файл имел то же имя, что и зависимость PyTorch, и менеджер пакетов PyPI (pip) установил этот вредоносный пакет вместо легитимного.

Эта атака на цепочку поставок, также известная как «dependency confusion», раскрыла чувствительную информацию Linux-машин с затронутыми версиями PyTorch-nightly, установленными через pip. 30 декабря 2022 года PyTorch объявил об инциденте и первых шагах по снижению риска, включая переименование и удаление зависимостей `torchtriton`.
