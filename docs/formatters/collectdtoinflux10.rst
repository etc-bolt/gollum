CollectdToInflux10
==================

CollectdToInflux10 provides a transformation from collectd JSON data to InfluxDB 0.9.1+ compatible line protocol data.
Trailing and leading commas are removed from the Collectd message beforehand.


Parameters
----------

**CollectdToInfluxFormatter**
  CollectdToInfluxFormatter defines the formatter applied before the conversion from Collectd to InfluxDB.
  By default this is set to format.Forward.

Example
-------

.. code-block:: yaml

	- "stream.Broadcast":
	    Formatter: "format.CollectdToInflux10"
	    CollectdToInflux10Formatter: "format.Forward"
