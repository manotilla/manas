# Manas Setup Documentation
Manas is running under systemd service and you can deploy `manas.service` file under the utils directory.
Deployment progress:

```sh
cp utils/manas.service /etc/system/systemd/manas.service
systemctl daemon-reload
systemctl start manas.service
```