(() => {
  // post element template
  const template = document.createElement('div');
  const postTime = document.createElement('span');
  const userName = document.createElement('strong');
  const postHeader = document.createElement('p')
  const post = document.createElement('p');
  const postContainer = document.createElement('div');
  const avatar = document.createElement('img');
  const avatarContainer = document.createElement('div');
  postTime.classList.add('post__date');
  userName.classList.add('post__user');
  postHeader.appendChild(userName);
  postHeader.appendChild(postTime);
  postHeader.classList.add('post__header');
  post.classList.add('post__main');
  postContainer.appendChild(postHeader);
  postContainer.appendChild(post);
  postContainer.classList.add('post__container');
  avatarContainer.appendChild(avatar);
  avatarContainer.classList.add('post__avatar');
  template.appendChild(avatarContainer);
  template.appendChild(postContainer);
  template.classList.add('post');

  document.addEventListener('DOMContentLoaded', () => {
    const host = document.querySelector('[data-host]').getAttribute('data-host');
    let socket = null;
    const chatBox = document.getElementById('chatBox');
    const msgBox = chatBox.querySelector('.chatTextArea');
    const messages = document.getElementById('messages');

    const createElm = (element, data) => {
      const date = new Date(data.When);
      const dateOpt = {
        weekday: "short", year: "numeric", month: "short",
        day: "numeric", hour: "2-digit", minute: "2-digit"
      };
      const elm = document.createElement(element);
      const content = template.cloneNode(true);
      content.querySelector('.post__date').textContent = ` - ${date.toLocaleTimeString('en-us', dateOpt)}`;
      content.querySelector('.post__user').textContent = data.Name;
      content.querySelector('.post__main').textContent = data.Message;
      content.querySelector('.post__avatar').children[0].src = data.AvatarURL || '/assets/images/user.png';
      elm.appendChild(content);
      return elm;
    }

    const submitHandler = e => {
      e.preventDefault();
      e.stopPropagation();
      if (!msgBox.value) {
        return false;
      }
      if (!socket) {
        alert('Error!!: cannot connect WebSocket.');
        return false;
      }
      socket.send(JSON.stringify({"Message": msgBox.value}));
      msgBox.value = '';
      return false;
    }

    if (!window['WebSocket']) {
      alert('Error!!: Cannot use WebSocket on your browser.\nPlease use modern browser like Google Chrome, Mozilla Firefox.');
    } else {
      socket = new WebSocket(`ws://${host}/room`);
      socket.onclose = () => {
        console.log('connecting closed.');
      }
      socket.onmessage = e => {
        const msg = JSON.parse(e.data);
        messages.appendChild(createElm('li',msg));
        messages.scrollTop = messages.scrollHeight;
      }
    }
    chatBox.addEventListener('submit', submitHandler);
    msgBox.addEventListener('keydown', e => {
      if (e.key === 'Enter' && e.shiftKey === true) {
        submitHandler(e)
      }
    })
  });
})();