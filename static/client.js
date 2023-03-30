let source = null;

const connectBtn = document.querySelector("#connect-btn");
const disconnectBtn = document.querySelector("#disconnect-btn");
const log = document.querySelector("#log");

connectBtn.addEventListener("click", () => {
  // Create the EventSource object and add the message event listener
  source = new EventSource("/stream");
  source.addEventListener("message", (event) => {
    // Append the message to the log element
    log.innerHTML += `${event.data}<br>`;
  });

  // Disable the connect button and enable the disconnect button
  connectBtn.disabled = true;
  disconnectBtn.disabled = false;
});

disconnectBtn.addEventListener("click", () => {
  // Close the EventSource connection
  source.close();

  // Enable the connect button and disable the disconnect button
  connectBtn.disabled = false;
  disconnectBtn.disabled = true;
});
