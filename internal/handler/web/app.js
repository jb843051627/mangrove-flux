const result = document.querySelector("#result");
  const id = encodeURIComponent(document.querySelector("#station").value);
  const response = await fetch(`/reports/${id}`);
  result.textContent = await response.text();
});
