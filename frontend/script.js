async function analyze() {
    const url = document.getElementById("urlInput").value;
    if (!url) {
        alert("Please enter a URL");
        return;
    }

    try {
        const response = await fetch("http://localhost:8080/analyze", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ url }),
        });

        if (!response.ok) {
            throw new Error(`Analyzing Failed.`);
        }

        const result = await response.json();
        displayResult(result); 
    } catch (error) {
        console.error("Error:", error);
        document.getElementById("output").textContent = `Error: ${error.message}`;
    }
}

function displayResult(data) {
    const outputDiv = document.getElementById("output");
    outputDiv.innerHTML = ''; 

    const title = document.createElement('h2');
    title.textContent = `Title: ${data.title}`;
    outputDiv.appendChild(title);

    const htmlVersion = document.createElement('p');
    htmlVersion.textContent = `HTML Version: ${data.html_version || 'Not specified'}`;
    outputDiv.appendChild(htmlVersion);

    const internalLinks = document.createElement('p');
    internalLinks.textContent = `Internal Links: ${data.internal_links}`;
    outputDiv.appendChild(internalLinks);

    const externalLinks = document.createElement('p');
    externalLinks.textContent = `External Links: ${data.external_links}`;
    outputDiv.appendChild(externalLinks);

    const inaccessibleLinks = document.createElement('p');
    inaccessibleLinks.textContent = `Inaccessible Links: ${data.inaccessible_links}`;
    outputDiv.appendChild(inaccessibleLinks);

    const loginForm = document.createElement('p');
    loginForm.textContent = `Login Form Present: ${data.login_form ? 'Yes' : 'No'}`;
    outputDiv.appendChild(loginForm);
}

function clearOutput() {
    document.getElementById("urlInput").value = ''; 
    document.getElementById("output").textContent = '';
}