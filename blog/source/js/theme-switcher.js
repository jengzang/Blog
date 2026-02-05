// 主題風格切換器
// Theme Style Switcher

const themeStyles = {
  default: {
    name: '默認 Default',
    primary: '#49b1f5',
    background: '#ffffff'
  },
  tech: {
    name: '技術感 Tech',
    primary: '#00ff00',
    background: '#0a0a0a'
  },
  academic: {
    name: '學術風 Academic',
    primary: '#2c3e50',
    background: '#ecf0f1'
  },
  warm: {
    name: '溫暖 Warm',
    primary: '#e74c3c',
    background: '#fff5f5'
  }
};

// 初始化主題
function initThemeStyleSwitcher() {
  const savedStyle = localStorage.getItem('theme-style') || 'default';
  applyThemeStyle(savedStyle);

  // 創建切換按鈕
  createStyleSwitcher();
}

// 應用主題風格
function applyThemeStyle(styleName) {
  const style = themeStyles[styleName];
  if (!style) return;

  document.documentElement.style.setProperty('--theme-primary', style.primary);
  document.documentElement.style.setProperty('--theme-bg', style.background);

  localStorage.setItem('theme-style', styleName);
}

// 創建切換器 UI
function createStyleSwitcher() {
  const switcher = document.createElement('div');
  switcher.className = 'theme-style-switcher';
  switcher.innerHTML = `
    <button class="style-btn" onclick="toggleStyleMenu()">
      <i class="fas fa-palette"></i>
    </button>
    <div class="style-menu" id="styleMenu" style="display: none;">
      ${Object.keys(themeStyles).map(key => `
        <div class="style-option" onclick="switchThemeStyle('${key}')">
          ${themeStyles[key].name}
        </div>
      `).join('')}
    </div>
  `;

  document.body.appendChild(switcher);
}

// 切換風格菜單
function toggleStyleMenu() {
  const menu = document.getElementById('styleMenu');
  menu.style.display = menu.style.display === 'none' ? 'block' : 'none';
}

// 切換主題風格
function switchThemeStyle(styleName) {
  applyThemeStyle(styleName);
  toggleStyleMenu();
}

// 頁面加載時初始化
if (document.readyState === 'loading') {
  document.addEventListener('DOMContentLoaded', initThemeStyleSwitcher);
} else {
  initThemeStyleSwitcher();
}
