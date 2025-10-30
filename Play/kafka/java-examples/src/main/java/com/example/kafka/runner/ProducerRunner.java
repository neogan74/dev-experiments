package com.example.kafka.runner;

import com.example.kafka.model.Message;
import com.example.kafka.producer.KafkaProducerService;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

@Component
public class ProducerRunner implements CommandLineRunner {

    private final KafkaProducerService producerService;

    public ProducerRunner(KafkaProducerService producerService) {
        this.producerService = producerService;
    }

    @Override
    public void run(String... args) throws Exception {
        // Uncomment to send messages on startup
        /*
        System.out.println("📤 Starting to produce messages...\n");

        for (int i = 0; i < 10; i++) {
            Message message = new Message(
                    i,
                    "Hello Kafka message #" + i,
                    System.currentTimeMillis()
            );

            producerService.sendMessage(message);
            Thread.sleep(500);
        }

        System.out.println("✨ All messages sent!");
        */
    }
}
