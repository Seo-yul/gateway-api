document.addEventListener('DOMContentLoaded', function() {
    // 현재 URL에서 언어 코드 확인
    const isKorean = window.location.pathname.startsWith('/ko/');
    
    // body에 언어 클래스 추가
    document.body.classList.add(isKorean ? 'is-korean' : 'is-english');
    
    // 기본 언어 선택기 텍스트 변경
    updateLanguageSelector(isKorean);
  });
  
  function updateLanguageSelector(isKorean) {
    // 언어 선택기 버튼 찾기
    const langButton = document.querySelector('.md-header__button[aria-label="Select language"]');
    if (!langButton) return;
    
    // 기존 아이콘 대신 텍스트 표시
    const icon = langButton.querySelector('svg');
    if (icon) {
      // 기존 아이콘 숨기기
      icon.style.display = 'none';
      
      // 언어 텍스트 추가
      const langText = document.createElement('span');
      langText.textContent = isKorean ? '한국어' : 'English';
      langText.style.marginLeft = '5px';
      langButton.appendChild(langText);
    }
  }
  
  // 현재 URL 경로에서 언어 코드 확인
  function getCurrentLanguage() {
    const path = window.location.pathname;
    if (path.startsWith('/ko/')) {
      return 'ko';
    }
    return 'en';
  }
  
  // 언어 전환 함수
  function switchLanguage(lang) {
    const currentPath = window.location.pathname;
    let newPath;
    
    if (lang === 'ko') {
      // 영어에서 한국어로 전환
      if (currentPath.startsWith('/ko/')) {
        return; // 이미 한국어 페이지
      }
      newPath = '/ko' + currentPath;
    } else {
      // 한국어에서 영어로 전환
      if (currentPath.startsWith('/ko/')) {
        newPath = currentPath.substring(3); // '/ko/' 제거
      } else {
        return; // 이미 영어 페이지
      }
    }
    
    // 쿠키에 언어 설정 저장
    document.cookie = `preferred_language=${lang}; path=/; max-age=31536000`; // 1년간 유효
    
    // 새 URL로 이동
    window.location.href = newPath;
  }
  
  // 언어 선택기 초기화
  function initLanguageSelector() {
    // 언어 선택기 DOM 요소 생성
    const selector = document.querySelector('.language-selector');
    if (!selector) return;
    
    const button = document.createElement('button');
    button.className = 'language-dropdown-btn';
    
    const currentLang = getCurrentLanguage();
    button.innerHTML = currentLang === 'ko' ? '한국어' : 'English';
    
    const dropdown = document.createElement('div');
    dropdown.className = 'language-dropdown-content';
    
    const enOption = document.createElement('a');
    enOption.href = '#';
    enOption.textContent = 'English';
    enOption.onclick = function(e) {
      e.preventDefault();
      switchLanguage('en');
    };
    
    const koOption = document.createElement('a');
    koOption.href = '#';
    koOption.textContent = '한국어';
    koOption.onclick = function(e) {
      e.preventDefault();
      switchLanguage('ko');
    };
    
    dropdown.appendChild(enOption);
    dropdown.appendChild(koOption);
    selector.appendChild(button);
    selector.appendChild(dropdown);
    
    // body에 언어 클래스 추가
    document.body.classList.add(currentLang === 'ko' ? 'is-korean' : 'is-english');
  } 