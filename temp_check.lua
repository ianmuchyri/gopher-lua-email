-- Function to check temperature and send an email alert
function check_temperature(temp)
  local recipient = "xxxxxxxxxxxx"  -- Replace with the actual recipient email

  if temp > 30 then
      local subject = string.format("🔥 Warning: Temperature too high (%.2f°C)", temp)
      local body = "The temperature has exceeded 30°C. Please take necessary precautions."
      local result = send_email(recipient, subject, body)
      print(result)
  elseif temp < 15 then
      local subject = string.format("❄️ Warning: Temperature too low (%.2f°C)", temp)
      local body = "The temperature has dropped below 15°C. Please take necessary precautions."
      local result = send_email(recipient, subject, body)
      print(result)
  else
      print(string.format("✅ Temperature (%.2f°C) is within the normal range.", temp))
  end
end

-- Example temperature reading (Replace this with a real sensor input)
local temperature = 32  -- Change this value to test
check_temperature(temperature)
