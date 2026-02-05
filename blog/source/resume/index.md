---
title: 簡歷 | Resume | CV
date: 2026-02-05 16:43:15
---

<div id="auth-container">
  <div id="auth-form" style="display: block;">
    <h2>🔒 此頁面需要密鑰訪問 | Authentication Required</h2>
    <p>請輸入訪問密鑰 | Please enter the access key:</p>
    <input type="password" id="access-key" placeholder="Access Key" style="padding: 10px; width: 300px; margin: 10px 0;">
    <br>
    <button onclick="authenticate()" style="padding: 10px 20px; cursor: pointer;">驗證 | Verify</button>
    <p id="error-message" style="color: red; display: none;"></p>
  </div>

  <div id="resume-content" style="display: none;">
    <h2>個人簡歷 | Personal Resume</h2>

    <h3>基本信息 | Basic Information</h3>
    <ul>
      <li><strong>姓名 | Name:</strong> Your Name</li>
      <li><strong>研究方向 | Research Focus:</strong> 語言學、地理學 | Linguistics, Geography</li>
    </ul>

    <h3>教育背景 | Education</h3>
    <ul>
      <li><strong>博士 | Ph.D.:</strong> [University Name] - [Field] (Year)</li>
      <li><strong>碩士 | Master:</strong> [University Name] - [Field] (Year)</li>
      <li><strong>學士 | Bachelor:</strong> [University Name] - [Field] (Year)</li>
    </ul>

    <h3>研究方向 | Research Interests</h3>
    <h4>語言學 | Linguistics</h4>
    <ul>
      <li>理論語言學 | Theoretical Linguistics</li>
      <li>應用語言學 | Applied Linguistics</li>
      <li>語言習得 | Language Acquisition</li>
      <li>方言研究 | Dialectology</li>
    </ul>

    <h4>地理學 | Geography</h4>
    <ul>
      <li>人文地理 | Human Geography</li>
      <li>自然地理 | Physical Geography</li>
      <li>地理信息系統 | GIS</li>
      <li>區域研究 | Regional Studies</li>
    </ul>

    <h3>學術成果 | Academic Achievements</h3>
    <ul>
      <li>發表論文 X 篇 | Published X papers</li>
      <li>參與項目 Y 個 | Participated in Y projects</li>
      <li>獲獎記錄 | Awards and Honors</li>
    </ul>

    <h3>技能專長 | Skills</h3>
    <ul>
      <li>語言能力 | Languages: 中文、英文 | Chinese, English</li>
      <li>研究工具 | Research Tools: GIS軟件、統計分析 | GIS Software, Statistical Analysis</li>
      <li>其他技能 | Other Skills: 數據分析、田野調查 | Data Analysis, Field Research</li>
    </ul>

    <h3>聯繫方式 | Contact Information</h3>
    <ul>
      <li><strong>Email:</strong> your.email@example.com</li>
      <li><strong>知乎 | Zhihu:</strong> @YourZhihuID</li>
    </ul>
  </div>
</div>

<script>
async function authenticate() {
  const key = document.getElementById('access-key').value;
  const errorMsg = document.getElementById('error-message');

  try {
    const response = await fetch('/api/auth', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ key: key })
    });

    const data = await response.json();

    if (data.success) {
      document.getElementById('auth-form').style.display = 'none';
      document.getElementById('resume-content').style.display = 'block';
    } else {
      errorMsg.textContent = '密鑰錯誤，請重試 | Invalid key, please try again';
      errorMsg.style.display = 'block';
    }
  } catch (error) {
    errorMsg.textContent = '驗證失敗，請稍後重試 | Authentication failed, please try again later';
    errorMsg.style.display = 'block';
  }
}

// Check if already authenticated
async function checkAuth() {
  try {
    const response = await fetch('/api/verify');
    const data = await response.json();

    if (data.authenticated) {
      document.getElementById('auth-form').style.display = 'none';
      document.getElementById('resume-content').style.display = 'block';
    }
  } catch (error) {
    console.log('Not authenticated');
  }
}

// Check authentication on page load
checkAuth();
</script>
