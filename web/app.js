const API = "http://localhost:8080";

async function setKey() {
  const key = document.getElementById("set-key").value;
  const value = document.getElementById("set-value").value;
  const ttl = document.getElementById("set-ttl").value;

  await fetch(API + "/set", {
    method: "POST",
    headers: {"Content-Type": "application/json"},
    body: JSON.stringify({
      key,
      value,
      ttl: ttl ? parseInt(ttl) : 0
    })
  });

  alert("SET OK");
}

async function getKey() {
  const key = document.getElementById("get-key").value;
  const res = await fetch(API + "/get?key=" + key);
  const data = await res.json();

  document.getElementById("get-result").innerText =
    JSON.stringify(data, null, 2);
}

async function delKey() {
  const key = document.getElementById("del-key").value;
  const res = await fetch(API + "/del?key=" + key, { method: "DELETE" });
  const data = await res.json();
  
  if (data.deleted === false) {
    alert("Error: Key not found or could not be deleted");
  } else {
    alert("Deleted successfully");
  }
}

async function loadStats() {
  const res = await fetch(API + "/stats");
  const data = await res.json();
  document.getElementById("stats").innerText =
    JSON.stringify(data, null, 2);
}
