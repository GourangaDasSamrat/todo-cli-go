//  Theme Toggle (System-aware)

const themeToggle = document.getElementById("theme-toggle");
const html = document.documentElement;
const mediaQuery = window.matchMedia("(prefers-color-scheme: dark)");

// Toggle manually
themeToggle.addEventListener("click", () => {
  const isDark = html.classList.toggle("dark");
  localStorage.setItem("theme", isDark ? "dark" : "light");
});

// React to system changes ONLY if user didn't choose
mediaQuery.addEventListener("change", (e) => {
  if (!localStorage.getItem("theme")) {
    html.classList.toggle("dark", e.matches);
  }
});

//  Mobile Menu Toggle

const mobileMenuBtn = document.getElementById("mobile-menu-btn");
const sidebar = document.getElementById("sidebar");
const mobileOverlay = document.getElementById("mobile-overlay");

mobileMenuBtn.addEventListener("click", () => {
  sidebar.classList.toggle("-translate-x-full");
  mobileOverlay.classList.toggle("hidden");
});

mobileOverlay.addEventListener("click", () => {
  sidebar.classList.add("-translate-x-full");
  mobileOverlay.classList.add("hidden");
});

//  Active Navigation Link

const navLinks = document.querySelectorAll(".nav-link");
const sections = document.querySelectorAll("section[id]");

function updateActiveLink() {
  const scrollPosition = window.scrollY + 100;

  sections.forEach((section) => {
    const top = section.offsetTop;
    const height = section.offsetHeight;
    const id = section.id;

    if (scrollPosition >= top && scrollPosition < top + height) {
      navLinks.forEach((link) => {
        link.classList.toggle("active", link.getAttribute("href") === `#${id}`);
      });
    }
  });
}

window.addEventListener("scroll", updateActiveLink);
window.addEventListener("load", updateActiveLink);

//  Close Mobile Nav on Click

navLinks.forEach((link) => {
  link.addEventListener("click", () => {
    if (window.innerWidth < 1024) {
      sidebar.classList.add("-translate-x-full");
      mobileOverlay.classList.add("hidden");
    }
  });
});

window.addEventListener("scroll", updateActiveLink);
window.addEventListener("load", updateActiveLink);

//  Smooth Mobile Nav Close

navLinks.forEach((link) => {
  link.addEventListener("click", () => {
    if (window.innerWidth < 1024) {
      sidebar.classList.add("-translate-x-full");
      mobileOverlay.classList.add("hidden");
    }
  });
});
