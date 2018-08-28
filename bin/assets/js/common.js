document.addEventListener('DOMContentLoaded', () => {
  document.querySelector('.header__user').addEventListener('click', e => {
    e.preventDefault();
    e.stopPropagation();
    const userElm = document.querySelector('.header__user');
    if (userElm.classList.contains('isOpened')) {
      userElm.classList.remove('isOpened');
    } else {
      userElm.classList.add('isOpened');
    }
  })
  document.getElementById('app').addEventListener('click', () => {
    const userElm = document.querySelector('.header__user');
    if (userElm.classList.contains('isOpened')) {
      userElm.classList.remove('isOpened');
    }
  })
});
