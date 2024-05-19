from dynaconf import Dynaconf
from loguru import logger
import sys

settings = Dynaconf(
    envvar_prefix="DYNACONF",
    settings_files=["settings.toml", ".secrets.toml"],
)

logger.remove()
logger.add(sys.stdout, level=settings.LOG_LEVEL)

# `envvar_prefix` = export envvars with `export DYNACONF_FOO=bar`.
# `settings_files` = Load these files in the order.
