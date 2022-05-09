# Minimal Go Link Shortener

This is just a project for test purposes, not even close to be suitable for production.

`GET /`: List of shortened links

`POST /api/generate-link`: Form data => link

Example:
```html
<form method="POST" action="/api/generate-link">
    <input type="text" name="link">
    <button type="submit">Create</button>
</form>
```
