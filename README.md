# Gezondhied-hook

Gezondhied-hook is a plugin for [Gezondheid](https://github.com/LiamEderzeel/gezondheid) to add webhooks on failing requests.

## Usage

```yaml
- name: test.test
  url: https://test.test
  interval: 10s
  plugins:
    - name: "gezondheid-hook.so"
      config:
        method: "POST"
        url: "https://webhook.test"
        statusCodeMinimum: 100
```

