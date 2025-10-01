package com.example.kafka.producer;

import com.example.kafka.model.Message;
import com.fasterxml.jackson.core.JsonProcessingException;
import com.fasterxml.jackson.databind.ObjectMapper;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.kafka.core.KafkaTemplate;
import org.springframework.kafka.support.SendResult;
import org.springframework.stereotype.Service;

import java.util.concurrent.CompletableFuture;

@Service
public class KafkaProducerService {

    private final KafkaTemplate<String, String> kafkaTemplate;
    private final ObjectMapper objectMapper;

    @Value("${kafka.topic.name}")
    private String topicName;

    public KafkaProducerService(KafkaTemplate<String, String> kafkaTemplate, ObjectMapper objectMapper) {
        this.kafkaTemplate = kafkaTemplate;
        this.objectMapper = objectMapper;
    }

    public void sendMessage(Message message) {
        try {
            String jsonMessage = objectMapper.writeValueAsString(message);
            String key = String.valueOf(message.getId());

            CompletableFuture<SendResult<String, String>> future = kafkaTemplate.send(topicName, key, jsonMessage);

            future.whenComplete((result, ex) -> {
                if (ex == null) {
                    System.out.println("✅ Message sent successfully:");
                    System.out.println("   Topic: " + result.getRecordMetadata().topic());
                    System.out.println("   Partition: " + result.getRecordMetadata().partition());
                    System.out.println("   Offset: " + result.getRecordMetadata().offset());
                    System.out.println("   Message: " + message);
                } else {
                    System.err.println("❌ Failed to send message: " + ex.getMessage());
                }
            });

        } catch (JsonProcessingException e) {
            System.err.println("❌ Error serializing message: " + e.getMessage());
        }
    }
}
