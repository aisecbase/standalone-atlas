const filterInput = document.querySelector("[data-filter-input]");
const filterSelects = [...document.querySelectorAll("[data-filter-select]")];
const filterItems = [...document.querySelectorAll("[data-filter-item]")];
const sortButtons = [...document.querySelectorAll("[data-sort-key]")];
const matrixMaturityFilter = document.querySelector("[data-matrix-maturity-filter]");
const matrixMaturityCells = [...document.querySelectorAll("[data-matrix-maturity]")];
const matrixMaturityCount = document.querySelector("#matrixMaturityCount");
const matrixMaturityLabel = document.querySelector("[data-matrix-maturity-label]");
const matrixMaturityTicks = [...document.querySelectorAll("[data-matrix-maturity-tick]")];
const matrixPlatformFilters = [...document.querySelectorAll("[data-matrix-platform-filter]")];
const subtechToggle = document.querySelector("[data-subtech-toggle]");
const subtechToggleLabel = document.querySelector("[data-subtech-toggle-label]");
const matrixHeadingCounts = [...document.querySelectorAll("[data-matrix-heading-count]")];
const themeToggle = document.querySelector("[data-theme-toggle]");
const themeToggleLabel = document.querySelector("[data-theme-toggle-label]");
const initialSortButton = sortButtons.find((button) => button.classList.contains("is-active"));
const sortState = {
  key: initialSortButton?.dataset.sortKey || "",
  direction: initialSortButton?.dataset.sortDirection || "asc",
};

function setTheme(theme) {
  const nextTheme = theme === "dark" ? "dark" : "light";
  document.documentElement.dataset.theme = nextTheme;
  try {
    localStorage.setItem("atlas-theme", nextTheme);
  } catch (error) {
    // Theme persistence is optional when storage is unavailable.
  }
  if (themeToggle) {
    const isDark = nextTheme === "dark";
    themeToggle.setAttribute("aria-pressed", isDark ? "true" : "false");
    themeToggle.setAttribute("aria-label", isDark ? "Включить светлую тему" : "Включить тёмную тему");
  }
  if (themeToggleLabel) themeToggleLabel.textContent = nextTheme === "dark" ? "Тёмная" : "Светлая";
}

if (themeToggle) {
  setTheme(document.documentElement.dataset.theme);
  themeToggle.addEventListener("click", () => {
    setTheme(document.documentElement.dataset.theme === "dark" ? "light" : "dark");
  });
}

function applyFilters() {
  const query = filterInput ? filterInput.value.trim().toLowerCase() : "";
  for (const item of filterItems) {
    const text = (item.dataset.filterText || item.textContent).toLowerCase();
    const matchesQuery = !query || text.includes(query);
    const matchesSelects = filterSelects.every((select) => {
      const value = select.value.trim().toLowerCase();
      if (!value) return true;
      return (item.dataset[select.dataset.filterSelect] || "").toLowerCase().includes(value);
    });
    item.hidden = !(matchesQuery && matchesSelects);
  }
}

if (filterInput) filterInput.addEventListener("input", applyFilters);
for (const select of filterSelects) select.addEventListener("change", applyFilters);

function techniqueWord(count) {
  const mod10 = count % 10;
  const mod100 = count % 100;
  if (mod10 === 1 && mod100 !== 11) return "техника";
  if (mod10 >= 2 && mod10 <= 4 && (mod100 < 12 || mod100 > 14)) return "техники";
  return "техник";
}

function applyMatrixControls() {
  if (!matrixMaturityFilter) return;
  const maturityRanks = { feasible: 0, demonstrated: 1, realized: 2 };
  const maturityLabels = ["Возможная", "Продемонстрированная", "Реализованная"];
  const threshold = Number(matrixMaturityFilter.value || 0);
  const showSubtechniques = subtechToggle ? subtechToggle.getAttribute("aria-pressed") === "true" : true;
  const activePlatforms = new Set(
    matrixPlatformFilters
      .filter((button) => button.getAttribute("aria-pressed") === "true")
      .map((button) => button.dataset.matrixPlatformFilter)
  );
  const visibleIDs = new Set();
  const totalIDs = new Set();
  const parentVisibility = new Map();

  for (const cell of matrixMaturityCells) {
    if (cell.dataset.matrixId) totalIDs.add(cell.dataset.matrixId);
    const rank = maturityRanks[cell.dataset.matrixMaturity] ?? 0;
    const passesMaturity = rank >= threshold;
    const cellPlatforms = (cell.dataset.matrixPlatforms || "").split("||").filter(Boolean);
    const passesPlatform = !matrixPlatformFilters.length || cellPlatforms.some((platform) => activePlatforms.has(platform));
    const isSubtechnique = cell.hasAttribute("data-subtech-cell");
    const parentVisible = isSubtechnique ? parentVisibility.get(cell.dataset.subtechParent) !== false : true;
    const isVisible = passesMaturity && passesPlatform && (!isSubtechnique || (showSubtechniques && parentVisible));
    cell.hidden = !isVisible;
    if (isVisible && cell.dataset.matrixId) visibleIDs.add(cell.dataset.matrixId);
    if (!isSubtechnique && cell.dataset.matrixId) parentVisibility.set(cell.dataset.matrixId, isVisible);
  }
  if (matrixMaturityLabel) matrixMaturityLabel.textContent = maturityLabels[threshold] || maturityLabels[0];
  for (const tick of matrixMaturityTicks) {
    tick.classList.toggle("is-active", Number(tick.dataset.matrixMaturityTick) === threshold);
  }
  matrixMaturityFilter.setAttribute("aria-valuetext", maturityLabels[threshold] || maturityLabels[0]);
  if (subtechToggleLabel) subtechToggleLabel.textContent = showSubtechniques ? "Скрыть" : "Показать";
  if (subtechToggle) subtechToggle.setAttribute("aria-label", showSubtechniques ? "Скрыть подтехники" : "Показать подтехники");
  if (matrixMaturityCount) {
    matrixMaturityCount.textContent = `Показано: ${visibleIDs.size} из ${totalIDs.size} ${techniqueWord(totalIDs.size)}`;
  }
  for (const countNode of matrixHeadingCounts) {
    const index = [...countNode.closest(".matrix-head").children].indexOf(countNode.closest(".matrix-heading"));
    const matrixBody = countNode.closest(".matrix")?.querySelector(".matrix-body");
    const matrixColumn = matrixBody ? matrixBody.children[index] : null;
    if (!matrixColumn) continue;
    const visibleCount = [...matrixColumn.querySelectorAll("[data-matrix-maturity]")].filter((cell) => !cell.hidden).length;
    countNode.textContent = `${visibleCount} ${techniqueWord(visibleCount)}`;
  }
}

if (matrixMaturityFilter) {
  matrixMaturityFilter.addEventListener("input", applyMatrixControls);
  applyMatrixControls();
}

if (subtechToggle) {
  subtechToggle.addEventListener("click", () => {
    const isPressed = subtechToggle.getAttribute("aria-pressed") === "true";
    subtechToggle.setAttribute("aria-pressed", isPressed ? "false" : "true");
    applyMatrixControls();
  });
}

for (const button of matrixPlatformFilters) {
  button.addEventListener("click", () => {
    const isPressed = button.getAttribute("aria-pressed") === "true";
    button.setAttribute("aria-pressed", isPressed ? "false" : "true");
    applyMatrixControls();
  });
}

let openSelectField = null;

function closeSelectField(field) {
  if (!field) return;
  field.classList.remove("is-open");
  const button = field.querySelector("[data-select-button]");
  const menu = field.querySelector("[data-select-menu]");
  if (button) button.setAttribute("aria-expanded", "false");
  if (menu) menu.hidden = true;
  if (openSelectField === field) openSelectField = null;
}

function enhanceSelect(select, index) {
  if (select.dataset.enhancedSelect) return;
  select.dataset.enhancedSelect = "true";
  const field = document.createElement("div");
  field.className = "select-field";
  const button = document.createElement("button");
  const label = document.createElement("span");
  const menu = document.createElement("div");
  const menuID = `select-menu-${index}`;
  button.type = "button";
  button.className = "select-button";
  button.dataset.selectButton = "";
  button.setAttribute("aria-haspopup", "listbox");
  button.setAttribute("aria-expanded", "false");
  button.setAttribute("aria-controls", menuID);
  menu.className = "select-menu";
  menu.dataset.selectMenu = "";
  menu.id = menuID;
  menu.setAttribute("role", "listbox");
  menu.hidden = true;
  label.textContent = select.selectedOptions[0] ? select.selectedOptions[0].textContent : "";
  button.append(label);
  select.parentNode.insertBefore(field, select);
  field.append(select, button, menu);
  select.classList.add("is-enhanced-native");
  select.tabIndex = -1;
  select.setAttribute("aria-hidden", "true");

  function options() {
    return [...menu.querySelectorAll("[data-select-option]")];
  }
  function selectedOptionButton() {
    return options().find((option) => option.dataset.value === select.value) || options()[0];
  }
  function refresh() {
    const selected = select.selectedOptions[0];
    label.textContent = selected ? selected.textContent : "";
    for (const option of options()) {
      const active = option.dataset.value === select.value;
      option.classList.toggle("is-active", active);
      option.setAttribute("aria-selected", active ? "true" : "false");
    }
  }
  function open() {
    if (openSelectField && openSelectField !== field) closeSelectField(openSelectField);
    field.classList.add("is-open");
    button.setAttribute("aria-expanded", "true");
    menu.hidden = false;
    openSelectField = field;
  }
  function close() {
    closeSelectField(field);
  }
  function choose(value) {
    if (select.value !== value) {
      select.value = value;
      select.dispatchEvent(new Event("change", { bubbles: true }));
    }
    refresh();
    close();
    button.focus();
  }

  for (const nativeOption of select.options) {
    const option = document.createElement("button");
    option.type = "button";
    option.className = "select-option";
    option.dataset.selectOption = "";
    option.dataset.value = nativeOption.value;
    option.setAttribute("role", "option");
    option.textContent = nativeOption.textContent;
    option.addEventListener("click", () => choose(nativeOption.value));
    menu.append(option);
  }
  button.addEventListener("click", () => {
    field.classList.contains("is-open") ? close() : open();
  });
  button.addEventListener("keydown", (event) => {
    if (!["ArrowDown", "ArrowUp", "Enter", " "].includes(event.key)) return;
    event.preventDefault();
    open();
    const selected = selectedOptionButton();
    if (selected) selected.focus();
  });
  menu.addEventListener("keydown", (event) => {
    const items = options();
    const current = items.indexOf(document.activeElement);
    let next = current;
    if (event.key === "Escape") {
      event.preventDefault();
      close();
      button.focus();
      return;
    }
    if (event.key === "ArrowDown") next = Math.min(items.length - 1, current + 1);
    if (event.key === "ArrowUp") next = Math.max(0, current - 1);
    if (event.key === "Home") next = 0;
    if (event.key === "End") next = items.length - 1;
    if (next !== current) {
      event.preventDefault();
      items[next].focus();
    }
  });
  select.addEventListener("change", refresh);
  refresh();
}

document.addEventListener("click", (event) => {
  if (openSelectField && !openSelectField.contains(event.target)) closeSelectField(openSelectField);
});
document.addEventListener("keydown", (event) => {
  if (event.key === "Escape") closeSelectField(openSelectField);
});
document.querySelectorAll(".filter-select").forEach(enhanceSelect);

function sortValue(row, key) {
  const value = row.dataset[`sort${key.charAt(0).toUpperCase()}${key.slice(1)}`] || "";
  return value.trim().toLowerCase();
}

function applySort(key) {
  const button = sortButtons.find((item) => item.dataset.sortKey === key);
  const tbody = button ? button.closest("table").querySelector("tbody") : null;
  if (!tbody) return;
  const direction = sortState.key === key && sortState.direction === "asc" ? "desc" : "asc";
  sortState.key = key;
  sortState.direction = direction;
  const rows = [...tbody.querySelectorAll("[data-filter-item]")];
  rows.sort((a, b) => {
    const first = sortValue(a, key);
    const second = sortValue(b, key);
    if (!first && second) return 1;
    if (first && !second) return -1;
    const result = first.localeCompare(second, "ru", { numeric: true, sensitivity: "base" });
    return direction === "asc" ? result : -result;
  });
  tbody.append(...rows);
  for (const item of sortButtons) {
    const active = item.dataset.sortKey === key;
    item.classList.toggle("is-active", active);
    item.dataset.sortDirection = active ? direction : "";
    item.closest("th").setAttribute("aria-sort", active ? (direction === "asc" ? "ascending" : "descending") : "none");
  }
}

for (const button of sortButtons) button.addEventListener("click", () => applySort(button.dataset.sortKey));
