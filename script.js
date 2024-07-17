const apiUrl = "http://localhost:0606";

async function login() {
  const username = document.getElementById("username").value;
  const password = document.getElementById("password").value;

  const response = await fetch(`${apiUrl}/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ username, password }),
  });

  if (response.ok) {
    const data = await response.json();
    localStorage.setItem("token", data.token);
    fetchPosts();
  } else {
    alert("Login failed! Incorrect username or password.");
  }
}

async function fetchPosts() {
  const token = localStorage.getItem("token");
  const response = await fetch(`${apiUrl}/posts`, {
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  if (response.ok) {
    const data = await response.json();
    displayPosts(data.posts);
    document.getElementById("login-form").style.display = "none";
    document.getElementById("posts").style.display = "block";
  } else {
    alert("Error loading posts.");
  }
}

function displayPosts(posts) {
  const postList = document.getElementById("post-list");
  postList.innerHTML = "";

  posts.forEach((post) => {
    const postItem = document.createElement("div");
    postItem.className = "post-item";
    postItem.innerHTML = `
            <h3>${post.title}</h3>
            <p>${post.content}</p>
            <button onclick="deletePost(${post.id})">Delete</button>
        `;
    postList.appendChild(postItem);
  });
}

async function createPost() {
  const title = document.getElementById("post-title").value;
  const content = document.getElementById("post-content").value;
  const token = localStorage.getItem("token");

  const response = await fetch(`${apiUrl}/posts`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify({ title, content }),
  });

  if (response.ok) {
    fetchPosts(); // Refresh the posts
    document.getElementById("create-post-form").style.display = "none";
  } else {
    alert("Error creating post.");
  }
}

function showCreatePostForm() {
  document.getElementById("create-post-form").style.display = "block";
}

async function deletePost(id) {
  const token = localStorage.getItem("token");
  const response = await fetch(`${apiUrl}/posts/${id}`, {
    method: "DELETE",
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  if (response.ok) {
    fetchPosts(); // Refresh the posts
  } else {
    alert("Error deleting post.");
  }
}

function logout() {
  localStorage.removeItem("token");
  document.getElementById("login-form").style.display = "block";
  document.getElementById("posts").style.display = "none";
}
